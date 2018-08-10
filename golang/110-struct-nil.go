package main

import "fmt"

type A struct {
	ID int
}

func main() {
	a1 := &A{}
	if a1 == nil {
		fmt.Println("a1 is nil")
	}
	var a2 *A
	if a2 == nil {
		fmt.Println("a2 is nil")
	}
}
