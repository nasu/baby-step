package main

import "fmt"

func main() {
	var a interface{}
	a = "[OK]"
	fmt.Println(a == "[OK]")
}
