package httphelpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response helper
func Response(w http.ResponseWriter, status int, props JSONProps) {
	log.Printf("Response: %v", props)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(props)
}
