package goredis

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

func ExampleTransaction(addr string, looper func(name string, f func() error)) {
	client := redis.NewClient(&redis.Options{Addr: addr})
	defer client.Close()

	var transaction func(client *redis.Client, key string) error
	var retryCnt int
	transaction = func(client *redis.Client, key string) error {
		err := client.Watch(func(tx *redis.Tx) error {
			n, err := tx.Get(key).Int64()
			if err != nil && err != redis.Nil {
				return err
			}
			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				_ = pipe.Set(key, strconv.FormatInt(n+1, 10), 0)
				return nil
			})
			return err
		}, key)
		if err == redis.TxFailedErr {
			retryCnt++
			return transaction(client, key)
		}
		return err
	}

	key := "go-redis:counter"
	looper("redis", func() error { return transaction(client, key) })
	res, _ := client.Get(key).Int64()
	log.Println("redis-go:", res, "RetryCnt", retryCnt)
}
