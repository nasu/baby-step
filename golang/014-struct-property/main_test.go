package main

import "testing"

func BenchmarkHasCreatedAtWithInterface(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCreatedAtWithInterface(&S1{})
		hasCreatedAtWithInterface(&S2{})
	}
}

func BenchmarkHasCreatedAtWithReflect(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCreatedAtWithReflect(&S1{})
		hasCreatedAtWithReflect(&S2{})
	}
}
