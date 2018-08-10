package main

import "fmt"

func main() {
	a := map[int]int{1: 1, 2: 2, 3: 3}
	for i := range a {
		for j := range a {
			fmt.Println(i, j)
		}
	}
}
