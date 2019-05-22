package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jean-lopes/data-integration-challenge/pkg/models"
	"github.com/jean-lopes/data-integration-challenge/pkg/storage"
	uuid "github.com/satori/go.uuid"
)

func TestCreateCompanyHandler(t *testing.T) {
	mux := http.NewServeMux()
	store, _ := storage.OpenInMemoryStorage()

	id := uuid.NewV4()
	c := models.Company{Name: "UK", Zip: "00000"}
	c.ID = &id
	store.Save(c)

	mux.HandleFunc("/", CreateCompanyHandler(store))
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
			"Bad request (Invalid company)",
			`{}`,
			http.StatusTeapot,
		},
		{
			"Bad request (UK Violation)",
			`{"id":"60992c98-7228-424e-8727-0eafa9351054","name":"UK","zip":"00000"}`,
			http.StatusTeapot,
		},
		{
			"Normal",
			`{"name": "Teste", "zip": "12345"}`,
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
