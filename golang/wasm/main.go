package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello")
	node := js.Global().Get("document").Call("createTextNode", "Hello")
	js.Global().Get("document").Call("getElementById", "sandbox").Call("appendChild", node)
}
