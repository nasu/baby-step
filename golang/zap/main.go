package main

import (
	"log"
	"time"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func main() {
	//sugar = zap.NewExample().Sugar()
	logger, _ := zap.NewDevelopment()
	sugar = logger.Sugar()
	defer sugar.Sync()
	test1()
}
func test1() {
	test2()
}
func test2() {
	test3()
}
func test3() {
	sugar.Infow("zap.Sugar: Start main function", "time", time.Now().Local()) // stacktrace off
	sugar.Warnw("zap.Sugar: Start main function", "time", time.Now().Local()) // stacktrace on
	log.Printf("std.Log: Start main function: %s", time.Now().Local())
	time.Sleep(time.Second * 5)
}
