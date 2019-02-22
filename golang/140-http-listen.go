package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println(http.StatusText(500))
	log.Fatal(http.ListenAndServe("", nil)) // using :80
}
