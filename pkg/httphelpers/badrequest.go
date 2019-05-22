package httphelpers

import "net/http"

// BadRequest helper
func BadRequest(w http.ResponseWriter, errors []error) {
	Response(w, http.StatusTeapot, JSONProps{
		"important": "Relax and drink some 🍵 :)",
		"errors":    errors})
}
