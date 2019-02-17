package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}

	readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
		fmt.Println(count(ctx, ro))
		fmt.Println(ro.Timestamp())
		return nil
	}, Options{true})
	readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
		fmt.Println(count(ctx, ro))
		fmt.Println(ro.Timestamp())
		return nil
	}, Options{false})
}

type Options struct {
	useTimestampBound bool
}

type Querier interface {
	Query(ctx context.Context, stmt spanner.Statement) *spanner.RowIterator
}

func readOnlyTransaction(ctx context.Context, client *spanner.Client, f func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error, opts ...Options) error {
	var useTimestampBound bool
	for _, opt := range opts {
		if opt.useTimestampBound {
			useTimestampBound = true
		}
	}

	var ro *spanner.ReadOnlyTransaction
	if useTimestampBound {
		ro = client.ReadOnlyTransaction().WithTimestampBound(spanner.ExactStaleness(30 * time.Second))
	} else {
		ro = client.ReadOnlyTransaction()
	}
	defer ro.Close()
	err := f(ctx, ro)
	if err != nil {
		return err
	}
	return nil
}

func count(ctx context.Context, q Querier) (int64, error) {
	stmt := spanner.Statement{SQL: `SELECT COUNT(*) FROM Singers WHERE FirstName="FirstName" AND LastName="LastName"`}
	iter := q.Query(ctx, stmt)
	defer iter.Stop()
	row, err := iter.Next()
	if err == iterator.Done {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	var cnt int64
	if err := row.Columns(&cnt); err != nil {
		return 0, err
	}
	return cnt, nil
}
