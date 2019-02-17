package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type Singer struct {
	SingerId  string
	FirstName string
	LastName  string
	CreatedAt time.Time
}

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}

	diff(ctx, client)
	nilInsert(ctx, client)
}

// time型にnilをインサートする方法
func nilInsert(ctx context.Context, client *spanner.Client) {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// nil を指定するとエラーになる
		// use T(nil), not nil
		// これは他の型の場合でも同様
		/*
			_, err := txn.Update(ctx, spanner.Statement{
				SQL: `INSERT INTO Singers (SingerId, FirstName, LastName, CreatedAt) VALUES (@SingerId, @FirstName, @LastName, @CreatedAt)`,
				Params: map[string]interface{}{
					"SingerId":  999,
					"FirstName": "sample",
					"LastName":  "sample",
					"CreatedAt": nil,
				},
			})
		*/
		/*
			_, err := txn.Update(ctx, spanner.Statement{
				SQL: `INSERT INTO Singers (SingerId, FirstName, LastName) VALUES (@SingerId, @FirstName, @LastName)`,
				Params: map[string]interface{}{
					"SingerId":  999,
					"FirstName": "sample",
					"LastName":  "sample",
				},
			})
		*/
		_, err := txn.Update(ctx, spanner.Statement{
			SQL: `INSERT INTO Singers (SingerId, FirstName, LastName, CreatedAt) VALUES (@SingerId, @FirstName, @LastName, @CreatedAt)`,
			Params: map[string]interface{}{
				"SingerId":  998,
				"FirstName": "sample",
				"LastName":  "sample",
				"CreatedAt": spanner.NullTime{},
			},
		})

		return err
	})
	if err != nil {
		panic(err)
	}
}

// time型をクエリに渡して比較できるかどうか
func diff(ctx context.Context, client *spanner.Client) {
	ro := client.ReadOnlyTransaction()
	now := time.Now().UTC()
	fmt.Println("TIME:", now)
	iter := ro.Query(ctx, spanner.Statement{
		SQL: `SELECT SingerId, FirstName, LastName, CreatedAt FROM Singers WHERE CreatedAt > @Time`,
		Params: map[string]interface{}{
			// time型をそのままで比較可能
			"Time": now.Add(-1 * time.Hour * 3),
		},
	})
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		model := &Singer{}
		if err := row.Columns(&model.SingerId, &model.FirstName, &model.LastName, &model.CreatedAt); err != nil {
			panic(err)
		}
		log.Println(model)
	}
}
