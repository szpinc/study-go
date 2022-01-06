package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(response http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(response, "URL.Path:%q\n", request.URL.Path)
}

func counter(resp http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(resp, "Count:%d", count)
	mu.Unlock()
}
