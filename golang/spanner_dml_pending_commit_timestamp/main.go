package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"google.golang.org/api/iterator"
)

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	useMutation(ctx, sig)
	useDML(ctx, sig)
}

func useDML(ctx context.Context, sig string) {
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}
	ts, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		ids := [2]string{generateId(), generateId()}
		if err := insertWithDML(ctx, txn, ids); err != nil {
			return errors.Wrap(err, "first insertWithDML")
		}
		// Error: Column CreatedAt cannot be accessed because it, or its associated index, has a pending CommitTimestamp
		/*
		   if err := insertWithDML(ctx, txn); err != nil {
		       return errors.Wrap(err, "second insertWithDML")
		   }
		*/
		if err := updateWithDML(ctx, txn, ids); err != nil {
			return errors.Wrap(err, "updateWithDML")
		}
		fmt.Println(time.Now(), "In transaction")
		time.Sleep(time.Second)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Confirming timestamp...")
	confirm(ctx, client, ts)
}

func useMutation(ctx context.Context, sig string) {
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}
	ts, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		if err := insertWithMutation(ctx, txn); err != nil {
			return errors.Wrap(err, "first insertWithMutation")
		}
		if err := insertWithMutation(ctx, txn); err != nil {
			return errors.Wrap(err, "second insertWithMutation")
		}
		fmt.Println(time.Now(), "In transaction")
		time.Sleep(time.Second)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Confirming timestamp...")
	confirm(ctx, client, ts)
}

func confirm(ctx context.Context, client *spanner.Client, ts time.Time) {
	stmt := spanner.Statement{
		SQL: `SELECT Id, CreatedAt FROM PendingCommitTimestampTable
        WHERE CreatedAt = @Time
        `,
		Params: map[string]interface{}{
			"Time": ts,
		},
	}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		var id string
		var createdAt time.Time
		if err := row.Columns(&id, &createdAt); err != nil {
			panic(err)
		}
		fmt.Println(createdAt.Local(), id)
	}
	fmt.Println(ts)
}

func insertWithDML(ctx context.Context, txn *spanner.ReadWriteTransaction, ids [2]string) error {
	// DML を使う場合は、同一カラムで PENDING_COMMIT_TIMESTAMP()と自分で指定を併用はできない
	// PENDING_COMMIT_TIMESTAMP()を使うなら、そのクエリの中ではすべてのレコードに適用するか
	// 全てで使わないかしかできない
	// 下記でいうと1つ目のCreatedAtだけPENDING...として2つ目を@CreatedAtとして外から渡したり
	// CURRENT_TIMESTAMP() を使ったりはできない
	stmt := spanner.Statement{
		SQL: `INSERT PendingCommitTimestampTable (Id, Name, CreatedAt, UpdatedAt)
        VALUES
        (@Id1, @Name, PENDING_COMMIT_TIMESTAMP(), @UpdatedAt),
        (@Id2, @Name, PENDING_COMMIT_TIMESTAMP(), @UpdatedAt)
        `,
		Params: map[string]interface{}{
			"Id1":  ids[0],
			"Id2":  ids[1],
			"Name": "foobar",
			//"UpdatedAt": time.Now().Add(time.Hour * 24), // 未来は指定できない
			"UpdatedAt": time.Now().Add(-time.Hour * 24),
		},
	}
	if _, err := txn.Update(ctx, stmt); err != nil {
		return err
	}
	return nil
}

func updateWithDML(ctx context.Context, txn *spanner.ReadWriteTransaction, ids [2]string) error {
	var stmt spanner.Statement
	// 同一トランザクション内で同じテーブルに2回 PENDING_COMMIT_TIMESTAMP()を走らせることはできない
	/*
		for _, id := range ids {
			fmt.Println("UPDATE... id=", id)
			stmt = spanner.Statement{
				SQL: `UPDATE PendingCommitTimestampTable SET Name=@Name, UpdatedAt=PENDING_COMMIT_TIMESTAMP() WHERE Id=@Id`,
				Params: map[string]interface{}{
					"Id":   id,
					"Name": "barfoo",
				},
			}
			if _, err := txn.Update(ctx, stmt); err != nil {
				return err
			}
		}
	*/
	// INSERT で PENDING_COMMIT_TIMESTAMP() を使っているカラムに対しては同一トランザクション内でUPDATEできない
	/*
		stmt = spanner.Statement{
			SQL: `UPDATE PendingCommitTimestampTable SET Name=@Name, CreatedAt=PENDING_COMMIT_TIMESTAMP() WHERE 1=1`,
			Params: map[string]interface{}{
				"Name": "barfoo",
			},
		}
		if _, err := txn.Update(ctx, stmt); err != nil {
			return err
		}
	*/
	// 未来の時間を挿入はできない
	/*
		stmt = spanner.Statement{
			SQL: `UPDATE PendingCommitTimestampTable SET Name=@Name, UpdatedAt=@UpdatedAt WHERE 1=1`,
			Params: map[string]interface{}{
				"Name":      "barfoo",
				"UpdatedAt": time.Now().Add(time.Second),
			},
		}
		if _, err := txn.Update(ctx, stmt); err != nil {
			return err
		}
	*/
	// 事前に適当な値を詰めて、その後commit_timestamp()を使う分にはOK
	stmt = spanner.Statement{
		SQL: `UPDATE PendingCommitTimestampTable SET Name=@Name, UpdatedAt=@UpdatedAt WHERE 1=1`,
		Params: map[string]interface{}{
			"Name":      "barfoo",
			"UpdatedAt": time.Now().Add(-time.Second),
		},
	}
	if _, err := txn.Update(ctx, stmt); err != nil {
		return err
	}
	// INSERTでPENDING_COMMIT_TIMESTAMP()を使っていないカラムに対して まとめて UPDATE する分には可能
	stmt = spanner.Statement{
		SQL: `UPDATE PendingCommitTimestampTable SET Name=@Name, UpdatedAt=PENDING_COMMIT_TIMESTAMP() WHERE 1=1`,
		Params: map[string]interface{}{
			"Name": "barfoo",
		},
	}
	if _, err := txn.Update(ctx, stmt); err != nil {
		return err
	}
	return nil
}

func insertWithMutation(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
	cols := []string{`Id`, `Name`, `CreatedAt`, `UpdatedAt`}
	// Mutation を使う場合は、commit_timestampと自分で指定する場合を併用できる
	txn.BufferWrite([]*spanner.Mutation{
		spanner.Insert(`PendingCommitTimestampTable`, cols, []interface{}{generateId(), "barfoo", spanner.CommitTimestamp, spanner.CommitTimestamp}),
		spanner.Insert(`PendingCommitTimestampTable`, cols, []interface{}{generateId(), "barfoo", spanner.CommitTimestamp, time.Time{}}),
	})
	return nil
}

func generateId() string {
	v4, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return v4.String()
}
