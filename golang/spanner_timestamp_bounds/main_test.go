package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/spanner"
)

func initBenchmark() *spanner.Client {
	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[0:]
			break
		}
		args = args[1:]
	}
	ctx := context.Background()
	sig := fmt.Sprintf("projects/%s/instances/%s/databases/%s", args[1], args[2], args[3])
	client, err := spanner.NewClient(ctx, sig)
	if err != nil {
		panic(err)
	}
	return client
}

func BenchmarkStaleRead(b *testing.B) {
	ctx := context.Background()
	client := initBenchmark()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
			count(ctx, ro)
			return nil
		}, Options{true})
	}
}

func BenchmarkStrongRead(b *testing.B) {
	ctx := context.Background()
	client := initBenchmark()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		readOnlyTransaction(ctx, client, func(ctx context.Context, ro *spanner.ReadOnlyTransaction) error {
			count(ctx, ro)
			return nil
		}, Options{false})
	}
}
