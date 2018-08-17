package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/spanner"
	"go.opencensus.io/trace"
)

func main() {
	args := os.Args
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			s := time.Now().UnixNano()
			ctx, span := trace.StartSpan(ctx, "example.com/Run")
			defer span.End()
			client, err := spanner.NewClient(ctx, sig)
			if err != nil {
				panic(err)
			}
			defer client.Close()

			t, err := client.ReadWriteTransaction(ctx, handle)
			if err != nil {
				panic(err)
			}
			log.Println(t, "Delay:", float64(time.Now().UnixNano()-s)/(1000*1000))
		}()
	}
	wg.Wait()
}

func handle(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
	ms := make([]*spanner.Mutation, 0)
	for _, i := range []int64{1, 2, 3} {
		m, err := getRecord(ctx, txn, i)
		if err != nil {
			return err
		}
		ms = append(ms, m)
	}
	time.Sleep(50 * time.Millisecond)

	if err := txn.BufferWrite(ms); err != nil {
		//return errors.Wrap(err, "hoge") // wrapすると spannerが isAborted判定できなくなる
		return err
	}
	outputs := make([]spanner.Mutation, 0)
	for _, m := range ms {
		outputs = append(outputs, *m)
	}
	log.Println(outputs)
	//time.Sleep(1 * time.Second)
	return nil
}

func getRecord(ctx context.Context, txn *spanner.ReadWriteTransaction, id int64) (*spanner.Mutation, error) {
	row, err := txn.ReadRow(ctx, "Test", spanner.Key{id}, []string{"Cnt"})
	if err != nil {
		return nil, err
	}
	var cnt int64
	if err := row.Column(0, &cnt); err != nil {
		return nil, err
	}
	cnt++
	m := spanner.Update("Test", []string{"ID", "Name", "Cnt"}, []interface{}{id, "nasu", cnt})
	return m, nil
}
