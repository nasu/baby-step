package main

import "testing"

func BenchmarkReference(b *testing.B) {
	a := newS()
	for i := 0; i < b.N; i++ {
		reference(a)
	}
}

func BenchmarkBuiltinShallowCopy(b *testing.B) {
	a := newS()
	for i := 0; i < b.N; i++ {
		builtinShallowCopy(a)
	}
}

func BenchmarkSelfShallowCopy(b *testing.B) {
	a := newS()
	for i := 0; i < b.N; i++ {
		selfShallowCopy(a)
	}
}

func BenchmarkJinzhuCopier(b *testing.B) {
	a := newS()
	for i := 0; i < b.N; i++ {
		jinzhuCopier(a)
	}
}

func BenchmarkUluleDeepCoiper(b *testing.B) {
	a := newS()
	for i := 0; i < b.N; i++ {
		ululeDeepCopier(a)
	}
}

func BenchmarkMohaeDeepcopy(b *testing.B) {
	a := newS()
	for i := 0; i < b.N; i++ {
		mohaeDeepcopy(a)
	}
}

func BenchmarkSelfDeepcopy(b *testing.B) {
	a := newS()
	for i := 0; i < b.N; i++ {
		selfDeepcopy(a)
	}
}
