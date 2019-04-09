package main

import (
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
	_ "./statik"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		log.Fatal("require port")
	}
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	log.Fatal(http.ListenAndServe(args[1], nil))
}
