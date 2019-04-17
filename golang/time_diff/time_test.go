package time_diff

import (
	"testing"
	"time"
)

var t1 = time.Now()
var t2 = time.Now().Add(time.Second)

func BenchmarkAfter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		After(t1, t2)
	}
}
func BenchmarkEqualAfter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualAfter(t1, t2)
	}
}
func BenchmarkUnix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unix(t1, t2)
	}
}
func BenchmarkEqualUnix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualUnix(t1, t2)
	}
}
func BenchmarkUnixNano(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UnixNano(t1, t2)
	}
}
func BenchmarkEqualUnixNano(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualUnixNano(t1, t2)
	}
}
