package main

import (
	"fmt"
	"time"
)

func main() {
	s := time.Now()
	elapsed := time.Since(s)
	fmt.Printf("String: %s\n", elapsed.String())
	fmt.Printf("Seconds: %f\n", elapsed.Seconds())
	fmt.Printf("Nanoseconds: %d\n", elapsed.Nanoseconds())
}
