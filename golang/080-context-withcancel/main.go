package main

import (
	"context"
	"fmt"
	"math"
	"runtime"
	"time"
)

var mem runtime.MemStats

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	profile("start")
	go tree(ctx, 1)
	profile("end first tree")

	cancel()
	profile("end cancel")
	time.Sleep(time.Second * 2)
	profile("main end")
}

func tree(ctx context.Context, i int) {
	if 5 < math.Log10(float64(i)) {
		return
	}
	ctx2, cancel := context.WithCancel(ctx)
	go tree(ctx2, i*10+1)
	go tree(ctx2, i*10+2)
	defer cancel()
	a := make([]int, 0)
	for {
		a = append(a, 1)
		select {
		case <-ctx.Done():
			profile(fmt.Sprintf("end cancel (%s %6d)", ctx.Err(), i))
			return
		}
	}
}

func profile(mes string) {
	runtime.ReadMemStats(&mem)
	fmt.Println(mes, mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}
