package main

import "fmt"

func main() {
	var a int
	exec(func() { a = 989 })
	fmt.Println(a)
}

func exec(f func()) {
	f()
}
