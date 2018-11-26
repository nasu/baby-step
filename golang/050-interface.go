package main

import "fmt"

type I interface {
	Query(string) int
}

type C1 struct {
}

type C2 struct {
	C1
}

func (c *C1) Query(str string) int {
	return 10
}

func main() {
	fmt.Println(GetC1().Query("a"))
	fmt.Println(GetC2().Query("a"))
}

func GetC1() I {
	return &C1{}
}

func GetC2() I {
	return &C2{}
}
