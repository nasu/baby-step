package main

import "fmt"

type I interface {
	Harakiri()
}

type A struct {
	Id int
}

func (a *A) Harakiri()  { a = nil }
func (a *A) GetId() int { return a.Id }

func main() {
	a := &A{1}
	f(a)
	fmt.Println(a) // &{1}
	a.Harakiri()
	fmt.Println(a) // &{1}
	a = nil
	fmt.Println(a) // <nil>

	a = &A{1}
	harakiriP(a)
	fmt.Println(a) // &{1}
	harakiriPP(&a)
	fmt.Println(a) // <nil>

	var b I = &A{1}
	fmt.Println((b.(*A)).GetId()) // 1
	fP(&b)
	fmt.Println(b) // <nil>

	c := &A{1}
	d := I(c)
	fP(&d)
	fmt.Println(c) // &{1}
	fmt.Println(d) // <nil>
}

func f(i I) {
	i = I(nil)
}
func fP(i *I) {
	*i = nil
}
func harakiriP(a *A) {
	a = nil
}
func harakiriPP(a **A) {
	*a = nil
}
