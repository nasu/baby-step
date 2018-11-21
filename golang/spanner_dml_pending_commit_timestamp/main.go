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
		if err := insertWithDML(ctx, txn); err != nil {
			return errors.Wrap(err, "first insertWithDML")
		}
		// Error: Column CreatedAt cannot be accessed because it, or its associated index, has a pending CommitTimestamp
		/*
		   if err := insertWithDML(ctx, txn); err != nil {
		       return errors.Wrap(err, "second insertWithDML")
		   }
		*/
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

func insertWithDML(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
	stmt := spanner.Statement{
		SQL: `INSERT PendingCommitTimestampTable (Id, Name, CreatedAt, UpdatedAt)
        VALUES
        (@Id1, @Name, PENDING_COMMIT_TIMESTAMP(), PENDING_COMMIT_TIMESTAMP()),
        (@Id2, @Name, PENDING_COMMIT_TIMESTAMP(), PENDING_COMMIT_TIMESTAMP())
        `,
		Params: map[string]interface{}{
			"Id1":  generateId(),
			"Id2":  generateId(),
			"Name": "foobar",
		},
	}
	if _, err := txn.Update(ctx, stmt); err != nil {
		return err
	}
	return nil
}

func insertWithMutation(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
	cols := []string{`Id`, `Name`, `CreatedAt`, `UpdatedAt`}
	txn.BufferWrite([]*spanner.Mutation{
		spanner.Insert(`PendingCommitTimestampTable`, cols, []interface{}{generateId(), "barfoo", spanner.CommitTimestamp, spanner.CommitTimestamp}),
		spanner.Insert(`PendingCommitTimestampTable`, cols, []interface{}{generateId(), "barfoo", spanner.CommitTimestamp, spanner.CommitTimestamp}),
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
