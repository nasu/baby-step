package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/seehuhn/mt19937"
)

func main() {
	fmt.Println(randomString(15))
	fmt.Println(randomString(15))

	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())
	for i := 0; i < 10; i++ {
		fmt.Println(rng.Int63n(100))
	}
	fmt.Println(rng.Int63n(1))
	//fmt.Println(rng.Int63n(0)) // panic
}

func randomString(n int) string {
	l := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = l[rand.Intn(len(l))]
	}
	return string(b)
}
