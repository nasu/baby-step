package main

import (
	"context"
	"fmt"
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
	ro := client.ReadOnlyTransaction()
	query(ctx, ro)
	queryWithStats(ctx, ro)
	analyzeQuery(ctx, ro)
}

func query(ctx context.Context, ro *spanner.ReadOnlyTransaction) {
	iter := ro.Query(ctx, spanner.Statement{
		SQL: `SELECT count(*) FROM Singers`,
	})
	defer iter.Stop()
	//color.Println(color.Green("WithoutStats", color.In))
	iterate(iter)
}

func queryWithStats(ctx context.Context, ro *spanner.ReadOnlyTransaction) {
	iter := ro.QueryWithStats(ctx, spanner.Statement{
		SQL: `SELECT count(*) FROM Singers`,
	})
	defer iter.Stop()
	//color.Println(color.Green("WithStats", color.In))
	iterate(iter)
}

func analyzeQuery(ctx context.Context, ro *spanner.ReadOnlyTransaction) {
	_, err := ro.AnalyzeQuery(ctx, spanner.Statement{
		SQL: `SELECT count(*) FROM Singers`,
	})
	if err != nil {
		panic(err)
	}
	/*
	   color.Println(color.Green("Analyze", color.In))
	   color.Println(color.Green("QueryPlan"))
	   fmt.Println(plan)
	*/
}

func iterate(iter *spanner.RowIterator) {
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
			break
		}

		var cnt int64
		if err := row.Columns(&cnt); err != nil {
			panic(err)
			break
		}
		//color.Println(color.Green("Result"))
		//fmt.Println(cnt)
	}
	/*
	   color.Println(color.Green("QueryPlan"))
	   p, err := json.MarshalIndent(iter.QueryPlan, "", "\t")
	   if err != nil {
	       panic(err)
	   }
	   fmt.Println(string(p))
	   color.Println(color.Green("QueryStats"))
	   for k, v := range iter.QueryStats {
	       fmt.Println("\t", k, "=>", v)
	   }
	*/
}
