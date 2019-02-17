package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
	single(ctx, client)
}

func single(ctx context.Context, client *spanner.Client) {
	ro := client.Single()
	cnt, err := count(ctx, ro)
	if err != nil {
		log.Println(err)
	}
	log.Println("Count 1:", cnt)
	cnt, err = count(ctx, ro)
	if err != nil {
		log.Println(err)
	}
	log.Println("Count 2:", cnt)
}

type Querier interface {
	Query(ctx context.Context, stmt spanner.Statement) *spanner.RowIterator
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
