package main

import (
	"context"
	"os"
	"testing"
)

/*
NOTE:
通信が計測のほとんど締めてしまうのでタイミングによってかなりばらつきがある
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

func BenchmarkSelectWithArrayStruct(b *testing.B) {
	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[0:]
			break
		}
		args = args[1:]
	}
	client := prepare(args)

	ctx := context.Background()
	ro := client.ReadOnlyTransaction()
	defer ro.Close()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SelectWithArrayStruct(ctx, ro)
	}
}

func BenchmarkSelectFlatten(b *testing.B) {
	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[0:]
			break
		}
		args = args[1:]
	}
	client := prepare(args)

	ctx := context.Background()
	ro := client.ReadOnlyTransaction()
	defer ro.Close()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SelectFlatten(ctx, ro)
	}
}
