package main

/*
BenchmarkStringDiff-8                   300000000                4.25 ns/op            0 B/op          0 allocs/op
BenchmarkStringSame-8                   200000000                7.28 ns/op            0 B/op          0 allocs/op
BenchmarkStringLongDiff-8               500000000                3.80 ns/op            0 B/op          0 allocs/op
BenchmarkStringLongSame-8                    100          10825261 ns/op               0 B/op          0 allocs/op
BenchmarkBytesEqualDiff-8               200000000                7.16 ns/op            0 B/op          0 allocs/op
BenchmarkBytesEqualSame-8               100000000               10.1 ns/op             0 B/op          0 allocs/op
BenchmarkBytesEqualLongDiff-8           200000000                6.88 ns/op            0 B/op          0 allocs/op
BenchmarkBytesEqualLongSame-8                100          10440526 ns/op               0 B/op          0 allocs/op

stringでコンバートしたほうが良い
*/

import (
	"math/rand"
	"testing"
)

func gen(n int) []byte {
	bytes := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return b
}

func clone(a []byte) []byte {
	b := make([]byte, len(a))
	for i, c := range a {
		b[i] = c
	}
	return b
}

func BenchmarkStringDiff(b *testing.B) {
	b1 := gen(100)
	b2 := gen(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_string(b1, b2)
	}
}

func BenchmarkStringSame(b *testing.B) {
	b1 := gen(100)
	b2 := clone(b1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_string(b1, b2)
	}
}

func BenchmarkStringLongDiff(b *testing.B) {
	b1 := gen(100 * 1000 * 1000)
	b2 := gen(100 * 1000 * 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_string(b1, b2)
	}
}

func BenchmarkStringLongSame(b *testing.B) {
	b1 := gen(100 * 1000 * 1000)
	b2 := clone(b1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_string(b1, b2)
	}
}

func BenchmarkBytesEqualDiff(b *testing.B) {
	b1 := gen(100)
	b2 := gen(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_bytesEqual(b1, b2)
	}
}

func BenchmarkBytesEqualSame(b *testing.B) {
	b1 := gen(100)
	b2 := clone(b1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_bytesEqual(b1, b2)
	}
}

func BenchmarkBytesEqualLongDiff(b *testing.B) {
	b1 := gen(100 * 1000 * 1000)
	b2 := gen(100 * 1000 * 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_bytesEqual(b1, b2)
	}
}

func BenchmarkBytesEqualLongSame(b *testing.B) {
	b1 := gen(100 * 1000 * 1000)
	b2 := clone(b1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_bytesEqual(b1, b2)
	}
}
