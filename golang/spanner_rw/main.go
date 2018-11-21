package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/iterator"
)

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	useDML(ctx, sig)
}

func useDML(ctx context.Context, sig string) {
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}
	id1 := generateStr()
	name1 := generateStr()
	fmt.Println("INSERT")
	_, err = client.ReadWriteTransaction(ctx, func(ctxInTxn context.Context, txn *spanner.ReadWriteTransaction) error {
		err := insert(ctxInTxn, txn, id1, name1)
		if err != nil {
			return errors.Wrap(err, "insert")
		}
		fmt.Println("Cnt with Client in transaction.", queryCountByClientAndId(ctx, client, id1) == 0)
		// Error. Cloud Spanner does not support nested transactions
		//fmt.Println("Cnt in transaction.", queryCountByClientAndId(ctxInTxn, client, id1) == 0)
		fmt.Println("Cnt with Txn in transaction.", queryCountByTxnAndId(ctx, txn, id1) == 1)
		return nil
	})
	fmt.Println("Cnt out with Client transaction.", queryCountByClientAndId(ctx, client, id1) == 1)

	fmt.Println("UPDATE")
	name2 := generateStr()
	_, err = client.ReadWriteTransaction(ctx, func(ctxInTxn context.Context, txn *spanner.ReadWriteTransaction) error {
		err := update(ctxInTxn, txn, id1, name2)
		if err != nil {
			return errors.Wrap(err, "update")
		}
		fmt.Println("Cnt of oldName with Client in transaction.", queryCountByClientByName(ctx, client, name1) == 1)
		fmt.Println("Cnt of newName with Client in transaction.", queryCountByClientByName(ctx, client, name2) == 0)
		fmt.Println("Cnt of oldName with Txn in transaction.", queryCountByTxnAndName(ctx, txn, name1) == 0)
		fmt.Println("Cnt of newName with Txn in transaction.", queryCountByTxnAndName(ctx, txn, name2) == 1)
		return nil
	})
	fmt.Println("Cnt of oldName with Client out transaction.", queryCountByClientByName(ctx, client, name1) == 0)
	fmt.Println("Cnt of newName with Client out transaction.", queryCountByClientByName(ctx, client, name2) == 1)
}

func queryCountByTxnAndId(ctx context.Context, txn *spanner.ReadWriteTransaction, id string) int64 {
	stmt := spanner.Statement{
		SQL: `SELECT Id FROM PendingCommitTimestampTable WHERE Id = @Id `,
		Params: map[string]interface{}{
			"Id": id,
		},
	}
	return queryCountByTxn(ctx, txn, stmt)
}

func queryCountByTxnAndName(ctx context.Context, txn *spanner.ReadWriteTransaction, name string) int64 {
	stmt := spanner.Statement{
		SQL: `SELECT Id FROM PendingCommitTimestampTable WHERE Name = @Name `,
		Params: map[string]interface{}{
			"Name": name,
		},
	}
	return queryCountByTxn(ctx, txn, stmt)
}

func queryCountByTxn(ctx context.Context, txn *spanner.ReadWriteTransaction, stmt spanner.Statement) int64 {
	iter := txn.Query(ctx, stmt)
	defer iter.Stop()
	var cnt int64
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		cnt++
	}
	return cnt
}

func queryCountByClientAndId(ctx context.Context, client *spanner.Client, id string) int64 {
	stmt := spanner.Statement{
		SQL: `SELECT Id FROM PendingCommitTimestampTable WHERE Id = @Id `,
		Params: map[string]interface{}{
			"Id": id,
		},
	}
	return queryCountByClient(ctx, client, stmt)
}

func queryCountByClientByName(ctx context.Context, client *spanner.Client, name string) int64 {
	stmt := spanner.Statement{
		SQL: `SELECT Id FROM PendingCommitTimestampTable WHERE Name = @Name `,
		Params: map[string]interface{}{
			"Name": name,
		},
	}
	return queryCountByClient(ctx, client, stmt)
}

func queryCountByClient(ctx context.Context, client *spanner.Client, stmt spanner.Statement) int64 {
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	var cnt int64
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		cnt++
	}
	return cnt
}

func insert(ctx context.Context, txn *spanner.ReadWriteTransaction, id, name string) error {
	stmt := spanner.Statement{
		SQL: `INSERT PendingCommitTimestampTable (Id, Name) VALUES (@Id, @Name)`,
		Params: map[string]interface{}{
			"Id":   id,
			"Name": name,
		},
	}
	if _, err := txn.Update(ctx, stmt); err != nil {
		return err
	}
	return nil
}

func update(ctx context.Context, txn *spanner.ReadWriteTransaction, id, name string) error {
	stmt := spanner.Statement{
		SQL: `UPDATE PendingCommitTimestampTable SET Name = @Name WHERE Id = @Id`,
		Params: map[string]interface{}{
			"Id":   id,
			"Name": name,
		},
	}
	if _, err := txn.Update(ctx, stmt); err != nil {
		return err
	}
	return nil
}

func generateStr() string {
	v4, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return v4.String()
}
