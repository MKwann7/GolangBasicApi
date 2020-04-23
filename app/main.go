package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hello() string {
	return "{success: true}"
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Hello())
}

func main() {
	http.HandleFunc("/", healthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
