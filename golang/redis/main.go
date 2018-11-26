package main

import (
	"log"
	"os"
	"sync"

	"github.com/nasu/baby-step/golang/redis/goredis"
	"github.com/nasu/baby-step/golang/redis/redigo"
)

const (
	keyPrefix = "counter"
	tryCnt    = 10
)

func main() {
	args := os.Args
	addr := args[1]
	redigo.ExampleTransaction(addr, looper)
	goredis.ExampleTransaction(addr, looper)
}

func looper(name string, f func() error) {
	var wg sync.WaitGroup
	for i := 0; i < tryCnt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := f()
			if err != nil {
				log.Println(name, err)
			}
		}()
	}
	wg.Wait()
}
