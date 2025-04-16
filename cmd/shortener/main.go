package main

import (
	"net/http"

	"github.com/timurtheprogrammer777/ya-urlshortener.git/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.PostMainHandler)
	http.HandleFunc("/{id}", handlers.GetIDHandler)
	http.ListenAndServe(":8080", nil)
}
