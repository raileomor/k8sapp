package main

import (
	"fmt"
	"log"
	"net/http"
)

const version = "0.1"

func main() {

	log.Printf("Starting server with version [v%s] on port 8080...", version)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s", r.RemoteAddr)
	fmt.Fprintf(w, "Building Apps for K8s app says Hi!")
}
