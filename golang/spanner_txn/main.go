package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type Querier interface {
	Query(ctx context.Context, stmt spanner.Statement) *spanner.RowIterator
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}
	nestSample(ctx, client)
	complexSample(ctx, client)
}

func nestSample(ctx context.Context, client *spanner.Client) {
	log.Println("--- Nest Sample start ---")
	log.Println("RwRw")
	if err := nestRwRw(ctx, client); err != nil {
		// Error when starting the second transaction.
		// Message = "Cloud Spanner does not support nested transactions"
		log.Println("RwRw", err)
	}
	log.Println("RwRo")
	if err := nestRwRo(ctx, client); err != nil {
		// Error when executing a query.
		// Message = "Cloud Spanner does not support nested transactions"
		log.Println("RwRo", err)
	}
	log.Println("RoRw")
	if err := nestRoRw(ctx, client); err != nil {
		// OK.
		log.Println("RoRw", err)
	}
	log.Println("RoRo")
	if err := nestRoRo(ctx, client); err != nil {
		// OK.
		log.Println("RoRo", err)
	}
	log.Println("--- Nest Sample end ---")
}

func complexSample(ctx context.Context, client *spanner.Client) {
	log.Println("--- Complex Sample start ---")
	var wg sync.WaitGroup
	roFirstCh := make(chan bool)
	rwCh := make(chan bool)
	wg.Add(2)
	go func(ctx context.Context, client *spanner.Client) {
		defer wg.Done()
		err := readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
			if cnt, err := count(ctx, ro); err == nil {
				log.Println("Count in RO:", cnt)
			} else {
				return err
			}
			// Wait for another transaction to insert.
			// We will see to the change.
			roFirstCh <- true
			log.Println("RO first end")
			<-rwCh
			err := readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
				if cnt, err := count(ctx, ro); err == nil {
					log.Println("Count in RO in RO:", cnt)
				} else {
					return err
				}
				log.Println("RO second end")
				return nil
			})
			if cnt, err := count(ctx, ro); err == nil {
				log.Println("Count in RO:", cnt)
			} else {
				return err
			}
			return err
		})
		if err != nil {
			log.Println("RO in RO Error:", err)
		}
	}(ctx, client)
	go func(ctx context.Context, client *spanner.Client) {
		<-roFirstCh
		defer func() {
			log.Println("RW end")
			rwCh <- true
			wg.Done()
		}()
		_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			return insert(ctx, txn)
		})
		if err != nil {
			log.Println("RW Error:", err)
		}
	}(ctx, client)
	wg.Wait()
	log.Println("--- Complex Sample end ---")
}

func insert(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
	stmt := spanner.Statement{
		SQL: `INSERT Singers (SingerId, FirstName, LastName) VALUES(@Id, @FirstName, @LastName)`,
		Params: map[string]interface{}{
			"Id":        time.Now().UnixNano(),
			"FirstName": "FirstName",
			"LastName":  "LastName",
		},
	}
	if _, err := txn.Update(ctx, stmt); err != nil {
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

func nestRwRw(ctx context.Context, client *spanner.Client) error {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			_, err := count(ctx, txn)
			return err
		})
		return err
	})
	return err
}

func nestRwRo(ctx context.Context, client *spanner.Client) error {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		err := readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
			_, err := count(ctx, ro)
			return err
		})
		return err
	})
	return err
}

func nestRoRw(ctx context.Context, client *spanner.Client) error {
	err := readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
		_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			_, err := count(ctx, txn)
			return err
		})
		return err
	})
	return err
}

func nestRoRo(ctx context.Context, client *spanner.Client) error {
	err := readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
		err := readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
			_, err := count(ctx, ro)
			return err
		})
		return err
	})
	return err
}

func readOnlyTransaction(ctx context.Context, client *spanner.Client, f func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error) error {
	ro := client.ReadOnlyTransaction()
	defer ro.Close()
	err := f(ctx, ro)
	if err != nil {
		return err
	}
	return nil
}
