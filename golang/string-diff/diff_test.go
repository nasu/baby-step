package diff

import (
	"math/rand"
	"testing"
	"time"
)

var s500, s5000, s50000, s500000 string

func init() {
	s500 = randomString(500)
	s5000 = randomString(5000)
	s50000 = randomString(50000)
	s500000 = randomString(500000)
}

func BenchmarkIsEmptyUsingLen500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingLen(s500)
	}
}

func BenchmarkIsEmptyUsingLen5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingLen(s5000)
	}
}

func BenchmarkIsEmptyUsingLen50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingLen(s50000)
	}
}

func BenchmarkIsEmptyUsingLen500000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingLen(s500000)
	}
}

func BenchmarkIsEmptyUsingEqual500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingEqual(s500)
	}
}

func BenchmarkIsEmptyUsingEqual5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingEqual(s5000)
	}
}

func BenchmarkIsEmptyUsingEqual50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingEqual(s50000)
	}
}

func BenchmarkIsEmptyUsingEqual500000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmptyUsingEqual(s500000)
	}
}

func randomString(num int) string {
	rand.Seed(time.Now().UnixNano())
	var l = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, num)
	for i := range b {
		b[i] = l[rand.Intn(len(l))]
	}
	return string(b)
}
