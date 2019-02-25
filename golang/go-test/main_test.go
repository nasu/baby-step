package sample

import (
	"log"
	"testing"
)

func TestSample(t *testing.T) {
	setup()
	teardown()
}

func setup() {
	log.Println("setup")
}

func teardown() {
	log.Println("teardown")
}
