package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "URL.Path:%q\n", request.URL.Path)
}
