package main

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/spanner"
)

/*
NOTE:
通信が計測のほとんど占めてしまうのでタイミングによってかなりばらつきがある
*/

// 一回 client作ってしまうとあとから実行したほうがキャッシュが効くせいか段違いで早くなってしまう
/*
var client *spanner.Client

func TestMain(m *testing.M) {
    args := os.Args
    for len(args) > 0 {
        if args[0] == "--" {
            args = args[0:]
            break
        }
        args = args[1:]
    }
    client = prepare(args)

    os.Exit(m.Run())
}
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
	client := prepare(args)
	return client.ReadOnlyTransaction()
}

func BenchmarkSelectWithArrayStruct(b *testing.B) {
	ctx := context.Background()
	ro := initBenchmark()
	defer ro.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SelectWithArrayStruct(ctx, ro)
	}
}

func BenchmarkSelectFlatten(b *testing.B) {
	ctx := context.Background()
	ro := initBenchmark()
	defer ro.Close()
	defer ro.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SelectFlatten(ctx, ro)
	}
}

func BenchmarkSelectWithArrayStructOnly(b *testing.B) {
	ctx := context.Background()
	ro := initBenchmark()
	defer ro.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SelectWithArrayStructOnly(ctx, ro)
	}
}

func BenchmarkSelectSimple(b *testing.B) {
	ctx := context.Background()
	ro := initBenchmark()
	defer ro.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SelectSimple(ctx, ro)
	}
}
