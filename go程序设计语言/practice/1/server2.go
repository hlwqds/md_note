package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int = 0
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %s\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	l_count := count
	mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", l_count)
}
