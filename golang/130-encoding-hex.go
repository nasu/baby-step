package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	a := []byte{0x61, 0x62, 0x63}
	fmt.Println(a)
	fmt.Println(string(a))
	fmt.Println(hex.EncodeToString(a))
}
