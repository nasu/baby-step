package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Atan2(0, 1) * 180 / math.Pi)
	fmt.Println(math.Atan2(1, 1) * 180 / math.Pi)
	fmt.Println(math.Atan2(1, 0) * 180 / math.Pi)
	fmt.Println(math.Atan2(1, -1) * 180 / math.Pi)
	fmt.Println(math.Atan2(0, -1) * 180 / math.Pi)
	fmt.Println(360 + math.Atan2(-1, -1)*180/math.Pi)
	fmt.Println(360 + math.Atan2(-1, 0)*180/math.Pi)
	fmt.Println(360 + math.Atan2(-1, 1)*180/math.Pi)
	fmt.Println(360 + math.Atan2(-0.000001, 1)*180/math.Pi)
}
