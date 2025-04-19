package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
)

var Store = make(map[string]string)

func PostMainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	url := string(body)
	id := generateID()

	Store[id] = url

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("http://localhost:8080/" + id))
}

func GetIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))

	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/")

	original, ok := Store[id]

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
