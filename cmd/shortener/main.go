package main

import (
	"fmt"
	"log"
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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("URL Shortener работает. Отправьте POST-запрос на /, чтобы сократить ссылку."))
	})
	fmt.Println("Server running on porn ", config.ServerConfig.Address)
	log.Fatal(http.ListenAndServe(config.ServerConfig.Address, r))

}
