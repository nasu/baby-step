package main

import "fmt"

type S1 struct {
	Id   int
	name string
	s2   *S2
}

type S2 struct {
	Id   int
	name string
}

func main() {
	//_ = S1{1} // too few values
	_ = S1{1, "", nil}
	_ = S1{Id: 1}

	s1 := S1{1, "s1 name", &S2{1, "s2 name"}}
	fmt.Println(s1.s2.name)
	fmt.Println(s1.s2.getName())
	fmt.Println(s1.s2.GetName())
}

func (s2 *S2) getName() string {
	return s2.name
}

func (s2 *S2) GetName() string {
	return s2.name
}
