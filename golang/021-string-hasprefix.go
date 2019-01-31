package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasPrefix("Gopher", "C"))  // false
	fmt.Println(strings.HasPrefix("Gopher", ""))   // true
}
