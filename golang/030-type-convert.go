package main

import "fmt"

type Byte16 [16]byte

func main() {
	var b16 Byte16 = Byte16{}
	//fmt.Println(string(b16)) // error
	fmt.Println(string(b16[:]))
}
