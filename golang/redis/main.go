package main

import (
	"log"
	"os"
	"sync"
	"time"

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
	s := time.Now()
	redigo.ExampleTransaction(addr, looper)
	log.Println("redigo:", time.Now().Sub(s))
	s = time.Now()
	goredis.ExampleTransaction(addr, looper)
	log.Println("goredis:", time.Now().Sub(s))
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
