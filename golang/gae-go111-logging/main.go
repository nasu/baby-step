package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/logging"
)

func main() {
	prj := os.Getenv("PROJECT")
	http.HandleFunc("/", index(prj))
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}

func index(prj string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := logging.NewClient(context.Background(), prj)
		if err != nil {
			return
		}
		defer client.Close()
		lg := client.Logger("mysample")
		js := []byte(`{"Name": "test", "Count": 3}`)
		lg.Log(logging.Entry{Payload: json.RawMessage(js), Severity: logging.Info})
		lg.Log(logging.Entry{Payload: json.RawMessage(js), Severity: logging.Warning})
		lg.StandardLogger(logging.Info).Println("hello")
		lg.StandardLogger(logging.Warning).Println("hello")
		log.Println("normal log")
		fmt.Println("normal log")
		longProcess(lg)
		fmt.Fprintf(w, "hello 2nd-gen")
	}
}

func longProcess(lg *logging.Logger) {
	for i := 0; i < 10; i++ {
		lg.StandardLogger(logging.Info).Println("count", i)
		time.Sleep(time.Second * 1)
	}
}
