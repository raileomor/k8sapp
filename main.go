package main

import (
	"fmt"
	"log"
	"net/http"
)

const version = "1.0"

func main() {

	log.Printf("[v%s]: Building Apps For K8s app", version)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Request received from %s", r.RemoteAddr)
	fmt.Fprintf(w, "[v%s] Building Apps For K8s app says Hi!!!\n", version)
}
