package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jean-lopes/data-integration-challenge/pkg/storage"
)

func TestMergeCompanyWebsiteHandler(t *testing.T) {
	mux := http.NewServeMux()
	store, _ := storage.OpenInMemoryStorage()
	mux.HandleFunc("/", MergeCompanyWebsiteHandler(store))
	tests := []struct {
		name string
		body string
		want int
	}{
		{
			"Bad request (Invalid JSON)",
			`{`,
			http.StatusTeapot,
		},
		{
			"Bad request (Empty json)",
			`{}`,
			http.StatusTeapot,
		},
		{
			"Bad request (invalid website)",
			`{"name": "Teste", "zip": "12345", "website": "sadpepe"}`,
			http.StatusTeapot,
		},
		{
			"Normal",
			`{"name": "Teste", "zip": "12345", "website": "https://www2.camara.leg.br/camaranoticias/tv/"}`,
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("POST", "/", strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)
			res := w.Result()
			if res.StatusCode != tt.want {
				t.Fatalf("Status code = %v, want %v", res.StatusCode, tt.want)
			}
		})
	}
}
