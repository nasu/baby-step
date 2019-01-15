package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nasu/baby-step/spanner_dml"

	"cloud.google.com/go/spanner"
)

type Singer struct {
	SingerID  string
	FirstName string
	LastName  string
}

var Singers = []Singer{
	Singer{
		SingerID:  "123456789012345678901234567890123456",
		FirstName: "Hikaru",
		LastName:  "Utada",
	},
	Singer{
		SingerID:  "123456789012345678901234567890123457",
		FirstName: "Yutaka",
		LastName:  "Ozaki",
	},
}

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.ReadWriteTransaction(ctx, insertUsingDML(ctx, client))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res:", res)

	err = insertUsingPartitionedDML()(ctx, client)
	if err != nil {
		fmt.Println(err)
	}
}

func insertUsingDML(ctx context.Context, client *spanner.Client) func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
	return func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		var sql string
		sql = fmt.Sprintf(`DELETE Singers WHERE true`)
		if err := spanner_dml.ExeUsingDML(ctx, txn, sql); err != nil {
			return err
		}
		for _, s := range Singers {
			sql = fmt.Sprintf(`INSERT Singers (SingerID, FirstName, LastName, LikeNum) VALUES('%s', '%s', '%s', %d)`, s.SingerID, s.FirstName, s.LastName, 0)
			if err := spanner_dml.ExeUsingDML(ctx, txn, sql); err != nil {
				return err
			}
		}
		// It's OK as *Not* Partitioned DML
		sql = fmt.Sprintf(`UPDATE Singers SET LikeNum = LikeNum + 1 WHERE true`)
		if err := spanner_dml.ExeUsingDML(ctx, txn, sql); err != nil {
			return err
		}
		return nil
	}
}

func insertUsingPartitionedDML() func(ctx context.Context, client *spanner.Client) error {
	return func(ctx context.Context, client *spanner.Client) error {
		var sql string
		sql = fmt.Sprintf(`DELETE Singers WHERE true`)
		if err := spanner_dml.ExeUsingPartitionedDML(ctx, client, sql); err != nil {
			return err
		}
		// Error - INSERT is not supported for Partitioned DML.
		for _, s := range Singers {
			sql = fmt.Sprintf(`INSERT Singers (SingerID, FirstName, LastName, LikeNum) VALUES('%s', '%s', '%s', %d)`, s.SingerID, s.FirstName, s.LastName, 0)
			if err := spanner_dml.ExeUsingPartitionedDML(ctx, client, sql); err != nil {
				fmt.Println("ERROR:", err)
			}
		}
		_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			for _, s := range Singers {
				sql = fmt.Sprintf(`INSERT Singers (SingerID, FirstName, LastName, LikeNum) VALUES('%s', '%s', '%s', %d)`, s.SingerID, s.FirstName, s.LastName, 0)
				if err := spanner_dml.ExeUsingDML(ctx, txn, sql); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return nil
		}
		// It's not good as not idempotent, but it can execute. Fmm...
		sql = fmt.Sprintf(`UPDATE Singers SET LikeNum = LikeNum + 1 WHERE true`)
		if err := spanner_dml.ExeUsingPartitionedDML(ctx, client, sql); err != nil {
			return err
		}

		return nil
	}
}
