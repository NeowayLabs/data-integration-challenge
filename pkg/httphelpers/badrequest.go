package httphelpers

import (
	"net/http"

	"github.com/jean-lopes/data-integration-challenge/pkg/util"
)

// BadRequest helper
func BadRequest(w http.ResponseWriter, errors []error) {

	Response(w, http.StatusTeapot, JSONProps{
		"important": "Relax and drink some 🍵 :)",
		"errors":    util.AsStrings(errors)})
}
