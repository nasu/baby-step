package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("", nil)) // using :80
}
