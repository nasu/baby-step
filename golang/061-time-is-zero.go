package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Time{}
	if t.IsZero() {
		fmt.Println("zero")
		fmt.Println("string", t)
		fmt.Println("epoch", t.Unix())
	}
}
