package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Funda!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting Funda server on :8080")
	http.ListenAndServe(":8080", nil)
}
