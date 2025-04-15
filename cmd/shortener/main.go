package main

import (
	"net/http"

	"github.com/timurtheprogrammer777/ya-urlshortener.git/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.PostMainHandler)
	http.HandleFunc("/{id}", handlers.GetIdHandler)
	http.ListenAndServe(":8081", nil)
}
