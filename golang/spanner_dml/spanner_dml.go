package spanner_dml

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
)

func ExeUsingDML(ctx context.Context, txn *spanner.ReadWriteTransaction, sql string) error {
	stmt := spanner.Statement{
		SQL: sql,
	}
	rowCount, err := txn.Update(ctx, stmt)
	if err != nil {
		return err
	}
	fmt.Println("SQL:", sql, "rowCount:", rowCount)
	return nil
}

func ExeUsingPartitionedDML(ctx context.Context, client *spanner.Client, sql string) error {
	stmt := spanner.Statement{
		SQL: sql,
	}
	rowCount, err := client.PartitionedUpdate(ctx, stmt)
	if err != nil {
		return err
	}
	fmt.Println("SQL:", sql, "rowCount:", rowCount)
	return nil
}
