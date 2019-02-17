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
