package main

import "fmt"

type Struct struct {
	Foo Foo
	Bar Bar
}

type Foo struct {
	Foo int
}

type Bar struct {
	Bar int
}

type Slice struct {
	Foo   Foo
	Bar   Bar
	Slice []int
}

func main() {
	a := make(map[Struct]bool)
	a[Struct{Foo{1}, Bar{1}}] = true
	a[Struct{Foo{2}, Bar{2}}] = true
	fmt.Println(a)

	// compile error when a struct includes slice.
	//b := make(map[Slice]bool)
}
