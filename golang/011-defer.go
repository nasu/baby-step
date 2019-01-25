package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	if true {
		fmt.Println("start block")
		defer fmt.Println("defer in block")
		fmt.Println("end block")
	}
	a()
	time.Sleep(time.Second)
	fmt.Println("end")
}

func a() {
	fmt.Println("start function")
	defer fmt.Println("defer in function")
	fmt.Println("end function")
}
