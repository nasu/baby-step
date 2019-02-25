package main

import "fmt"

func main() {
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println("append", a)
	a = nil
	fmt.Println("set nil", a)
	a = append(a, 1)
	fmt.Println("append", a)
	a = a[:0]
	fmt.Println("set a[:0]", a)
	a = append(a, 1)
	fmt.Println("append", a)
}
