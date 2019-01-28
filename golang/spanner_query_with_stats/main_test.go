package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/spanner"
)

/*
BenchmarkQuery-8                    1000          11531475 ns/op           16058 B/op        240 allocs/op
BenchmarkQueryWithStats-8           1000          14988492 ns/op           48471 B/op       1057 allocs/op
BenchmarkAnalyzeQuery-8             1000          11875962 ns/op           26507 B/op        454 allocs/op
*/

func initBenchmark() *spanner.ReadOnlyTransaction {
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
	return client.ReadOnlyTransaction()
}

func BenchmarkQuery(b *testing.B) {
	ctx := context.Background()
	ro := initBenchmark()
	defer ro.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		query(ctx, ro)
	}
}

func BenchmarkQueryWithStats(b *testing.B) {
	ctx := context.Background()
	ro := initBenchmark()
	defer ro.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		queryWithStats(ctx, ro)
	}
}

func BenchmarkAnalyzeQuery(b *testing.B) {
	ctx := context.Background()
	ro := initBenchmark()
	defer ro.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		analyzeQuery(ctx, ro)
	}
}
