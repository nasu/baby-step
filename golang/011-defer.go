package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start a")
	if true {
		fmt.Println("start block")
		defer fmt.Println("defer in block")
		fmt.Println("end block")
	}
	a()
	time.Sleep(time.Second)
	fmt.Println("end a")

	print("=====\n")
	fmt.Println("start b")
	fmt.Println(b())
	time.Sleep(time.Second)
	fmt.Println("end b")

	print("=====\n")
	fmt.Println("start c")
	fmt.Println(c())
	time.Sleep(time.Second)
	fmt.Println("end c")

}

func a() {
	fmt.Println("start function a")
	defer fmt.Println("defer in function a")
	fmt.Println("end function a")
}

func b() (err error) {
	defer func() {
		fmt.Println("defer in function b")
		err = errors.New("error")
	}()
	fmt.Println("end function b")
	return
}

func c() error {
	var err error
	defer func() error {
		fmt.Println("defer in function c")
		err = errors.New("error")
		return err
	}()
	fmt.Println("end function c")
	return err
}
