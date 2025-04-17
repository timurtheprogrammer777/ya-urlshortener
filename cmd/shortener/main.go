package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/timurtheprogrammer777/ya-urlshortener.git/internal/config"
	"github.com/timurtheprogrammer777/ya-urlshortener.git/internal/handlers"
)

func main() {
	config.ParseFlags()
	r := chi.NewRouter()

	r.Post("/", handlers.PostMainHandler)
	r.Get("/{id}", handlers.GetIDHandler)
	http.ListenAndServe(":8080", r)
}
