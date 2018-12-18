package spanner_dml

import (
	"context"
	"fmt"
	"os"
	"testing"

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

// これテストで走らせるとまともに動いてない。argsのいちがずれるので
func prepare() (context.Context, *spanner.Client) {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		fmt.Println(err)
	}
	client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		var sql string
		sql = fmt.Sprintf(`DELETE Singers WHERE true`)
		if err := ExeUsingDML(ctx, txn, sql); err != nil {
			return err
		}
		for _, s := range Singers {
			sql = fmt.Sprintf(`INSERT Singers (SingerID, FirstName, LastName, LikeNum) VALUES('%s', '%s', '%s', %d)`, s.SingerID, s.FirstName, s.LastName, 0)
			if err := ExeUsingDML(ctx, txn, sql); err != nil {
				return err
			}
		}
		return nil
	})
	return ctx, client
}

func BenchmarkExeUsingDML(b *testing.B) {
	ctx, client := prepare()
	sql := fmt.Sprintf(`UPDATE Singers SET LikeNum = LikeNum + 1 WHERE true`)
	for n := 0; n < b.N; n++ {
		client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			ExeUsingDML(ctx, txn, sql)
			return nil
		})
	}
}

func BenchmarkExeUsingPartitionedDML(b *testing.B) {
	ctx, client := prepare()
	sql := fmt.Sprintf(`UPDATE Singers SET LikeNum = LikeNum + 1 WHERE true`)
	for n := 0; n < b.N; n++ {
		ExeUsingPartitionedDML(ctx, client, sql)
	}
}
