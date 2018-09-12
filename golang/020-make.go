package main

import "fmt"

func main() {
	a := make([]int, 0)
	b := []int{}
	fmt.Println(a, b)
	fmt.Println(&a, &b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))
	a = append(a, 1)
	b = append(b, 1)

}
