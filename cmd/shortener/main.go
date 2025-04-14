package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"sync"
)

var (
	store = make(map[string]string)
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/", postHandler)
	http.HandleFunc("/", getHandler)
	http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Path != "/" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	url := string(body)
	id := generateID()

	mu.Lock()
	store[id] = url
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("http://localhost:8080/" + id))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path == "/" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/")

	mu.Lock()
	original, ok := store[id]
	mu.Unlock()

	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, original, http.StatusTemporaryRedirect)
}

func generateID() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:8]
}
