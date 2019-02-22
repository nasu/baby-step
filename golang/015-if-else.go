package main

import "fmt"

func main() {
	fmt.Println(a())
}

func a() string {
	if true {
		return "true"
	} else {
		return "false"
	}
}
