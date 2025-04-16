package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/timurtheprogrammer777/ya-urlshortener.git/internal/handlers"
)

func TestPostMainHandler(t *testing.T) {

	req := httptest.NewRequest("POST", "/", strings.NewReader("https://example.com"))
	w := httptest.NewRecorder()

	handlers.PostMainHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("ожидался статус %d, получили %d", http.StatusCreated, resp.StatusCode)
	}

	if string(body) == "https://example.com" {
		t.Errorf("Новая ссылка не сгенерировалась")
	}

}

func TestGetIDHandler_GlobalStore(t *testing.T) {
	handlers.Store["abc123"] = "https://example.com"

	req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
	w := httptest.NewRecorder()

	handlers.GetIDHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("expected 307, got %d", resp.StatusCode)
	}
}
