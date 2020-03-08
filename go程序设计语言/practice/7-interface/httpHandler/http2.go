package main

import (
	"log"
	"net/http"
)

func main() {
	db := database{
		"shoes": 99,
		"rice":  1000,
	}
	mux := http.NewServeMux()

	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
