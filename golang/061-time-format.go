package main

import (
	"fmt"
	"time"
)

func main() {
	parse("15:04:05-0700", "21:00:00+0900") // 0000-01-01 21:00:00 +0900 JST
	parse("15:04:05-0700", "21:00:00Z")     // error
	parse("15:04:05Z0700", "21:00:00+0900") // 0000-01-01 21:00:00 +0900 JST
	parse("15:04:05Z0700", "21:00:00Z")     // 0000-01-01 21:00:00 +0000 UTC
	fmt.Println(time.Duration(time.Second * 0).String())
}

func parse(f, s string) {
	t, err := time.Parse(f, s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
}
