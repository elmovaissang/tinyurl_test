package main

import (
	"log"
	"net/http"
)

func main() {
	srv := newServer()
	log.Println("tinyurl listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", srv.routes()))
}
