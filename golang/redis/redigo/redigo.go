package redigo

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
)

func ExampleTransaction(addr string, looper func(name string, f func() error)) {
	pool := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 5 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}

	var transaction func(key string) error
	var retryCnt int
	transaction = func(key string) error {
		conn := pool.Get()
		if conn.Err() == redis.ErrPoolExhausted {
			retryCnt++
			//log.Println("redigo", redis.ErrPoolExhausted)
			time.Sleep(time.Millisecond * 10)
			transaction(key)
		}
		defer conn.Close()

		if _, err := redis.String(conn.Do("WATCH", key)); err != nil {
			return errors.Wrap(err, "watch")
		}
		n, err := redis.Int(conn.Do("GET", key))
		if err != nil && err != redis.ErrNil {
			return errors.Wrap(err, "get")
		}
		if err := conn.Send("MULTI"); err != nil {
			return errors.Wrap(err, "multi")
		}
		if err := conn.Send("SET", key, n+1); err != nil {
			return errors.Wrap(err, "set")
		}
		res, err := conn.Do("EXEC")
		if !checkExecResult(res) || err != nil {
			retryCnt++
			transaction(key)
		}
		return err
	}

	key := "redigo:counter"
	looper("redis", func() error { return transaction(key) })
	conn := pool.Get()
	defer conn.Close()
	res, _ := redis.String(conn.Do("GET", key))
	log.Println("redis:", res, "RetryCnt:", retryCnt)
}

func checkExecResult(res interface{}) bool {
	switch v := res.(type) {
	case string:
		return true
		/*
		   if v == "[OK]" {
		       return true
		   }
		*/
	case []interface{}:
		for _, vv := range v {
			return checkExecResult(vv)
		}
	default:
		return false
	}
	return false
}
