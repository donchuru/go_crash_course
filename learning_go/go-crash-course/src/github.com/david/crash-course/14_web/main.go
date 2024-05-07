// Closures are Anonymous functions

package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from Go crash course")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}