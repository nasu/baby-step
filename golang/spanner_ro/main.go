package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	uuid "github.com/satori/go.uuid"
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
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		recycleReadOnlyTransaction(ctx, client)
		defer wg.Done()
	}()
	go func() {
		readOnlyTransaction(ctx, client)
		defer wg.Done()
	}()
	go func() {
		readWriteTransaction(ctx, client)
		defer wg.Done()
	}()
	wg.Wait()
}

func generateId() string {
	v4, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return v4.String()
}

func recycleReadOnlyTransaction(ctx context.Context, client *spanner.Client) {
	ro := client.ReadOnlyTransaction()
	defer ro.Close()
	for i := 0; i < 10; i++ {
		log.Println("Recycle RO start.", i)
		stmt := spanner.Statement{SQL: `SELECT SingerId FROM Singers`}
		iter := ro.Query(ctx, stmt)
		defer iter.Stop()
		var cnt int64
		for {
			_, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Println(err)
				return
			}
			cnt++
		}
		log.Println("Recycle RO end.", i, "RowCount=", cnt)
		time.Sleep(time.Millisecond * 1000)
	}
}

func readOnlyTransaction(ctx context.Context, client *spanner.Client) {
	for i := 0; i < 10; i++ {
		log.Println("RO start.", i)
		ro := client.ReadOnlyTransaction()
		defer ro.Close()
		stmt := spanner.Statement{SQL: `SELECT SingerId FROM Singers`}
		iter := ro.Query(ctx, stmt)
		defer iter.Stop()
		var cnt int64
		for {
			_, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Println(err)
				return
			}
			cnt++
		}
		log.Println("RO end.", i, "RowCount=", cnt)
		time.Sleep(time.Millisecond * 1000)
	}
}

func readWriteTransaction(ctx context.Context, client *spanner.Client) {
	for i := 0; i < 10; i++ {
		log.Println("RW start.", i)
		stmt := spanner.Statement{
			SQL: `INSERT Singers (SingerId, FirstName, LastName, LikeNum) VALUES(@Id, @FirstName, @LastName, @LikeNum)`,
			Params: map[string]interface{}{
				"Id":        generateId(),
				"FirstName": "a",
				"LastName":  "b",
				"LikeNum":   0,
			},
		}
		ts, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			if _, err := txn.Update(ctx, stmt); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("RW end.", i, "Commit Timestamp", ts.Local())
		time.Sleep(time.Millisecond * 500)
	}
}
