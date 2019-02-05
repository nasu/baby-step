package main

import (
    "fmt"
    "time"
)

func main() {
    t1 := time.Now()
    t2 := t1.Add(-1 * time.Second)
    //fmt.Println(t2 < t1) // panic
    fmt.Println(t2.Unix() < t1.Unix())
    fmt.Println(t2.Before(t1))
    fmt.Println(t1.After(t2))
    fmt.Println(t2.Equal(t2))
}
