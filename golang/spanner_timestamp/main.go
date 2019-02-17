package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/spanner"
)

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}
	commit(ctx, client)
}

func commit(ctx context.Context, client *spanner.Client) {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `INSERT Singers (SingerId, CreatedAt) VALUES(@Id, @Time)`,
			Params: map[string]interface{}{
				"Id": time.Now().UnixNano(),
				// 型違うよと言われる
				//"Time": `PENDING_COMMIT_TIMESTAMP()`,

				// spanner.commit_timestamp() が実行されるが int64扱い？になってデータ入らない
				//"Time": spanner.CommitTimestamp,

				// これは上手くいくが単に 0 になる
				//"Time": time.Time{}.In(time.FixedZone("CommitTimestamp placeholder", 0xDB)),
			},
		}
		_, err := txn.Update(ctx, stmt)
		return err
	})
	if err != nil {
		log.Println("DML", err)
	}
	_, err = client.Apply(ctx, []*spanner.Mutation{
		spanner.InsertOrUpdate("Singers", []string{"SingerId", "CreatedAt"}, []interface{}{time.Now().UnixNano(), spanner.CommitTimestamp}),
	})
	if err != nil {
		log.Println("Mutation", err)
	}
}
