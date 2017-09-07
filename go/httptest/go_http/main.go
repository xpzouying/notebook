package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", indexHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
