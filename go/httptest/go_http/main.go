package main

import (
	"fmt"
	"log"
	"net/http"
)

type stringHandle string

func (sh stringHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(sh))
}

var (
	handleSets = []struct {
		path string
		resp string
	}{
		{"/zouying", "hello zouying"},
		{"/tim", "hello tim"},
		{"/china", "great china"},
	}
)

func newHandle() http.Handler {
	mux := http.NewServeMux()
	for _, h := range handleSets {
		mux.Handle(h.path, stringHandle(h.resp))
	}

	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", newHandle()))
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
