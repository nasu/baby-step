package main

import "fmt"

type Stringer interface {
	String() string
}

type ConcreteA struct {
	str string
}

func (c ConcreteA) String() string {
	return c.str
}

type ConcreteB struct {
	str string
}

func main() {
	fmt.Println(GetStringer(ConcreteA{"abc"}).String())
	// panic: interface conversion: main.ConcreteB is not main.Stringer: missing method String
	//fmt.Println(GetStringer(ConcreteB{"abc"}))

	fmt.Println(CheckAndGetStringer(ConcreteA{"abc"}).String())
	fmt.Println(CheckAndGetStringer(ConcreteB{"abc"})) // nil
}

func GetStringer(a interface{}) Stringer {
	return a.(Stringer)
}

func CheckAndGetStringer(a interface{}) Stringer {
	if s, ok := a.(Stringer); ok {
		return s
	}
	return nil
}
