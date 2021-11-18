package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index page")
}

func next_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Next page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/next/", next_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
