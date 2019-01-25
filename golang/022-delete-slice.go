package main

import "fmt"

func main() {
	fmt.Println(deleteSlice([]string{"a", "b", "c"}, "a"))
	fmt.Println(deleteSlice([]string{"a", "b", "c"}, "b"))
	fmt.Println(deleteSlice([]string{"a", "b", "c"}, "c"))
	fmt.Println(deleteSlice([]string{"a", "b", "c"}, "d"))
	/*
	   [b c]
	   [a c]
	   [a b]
	   [a b c]
	*/
}

func deleteSlice(slice []string, str string) []string {
	for i, s := range slice {
		if s == str {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
