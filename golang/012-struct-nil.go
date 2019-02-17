package main

import "fmt"

type S struct {
	I int
}

func (s *S) GetI() int {
	if s == nil {
		return -1
	}
	return s.I
}

func main() {
	var s *S
	fmt.Println(s.GetI())
	//fmt.Println(s.I) // panic
}
