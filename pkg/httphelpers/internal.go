package httphelpers

import (
	"net/http"
)

// InternalError helper
func InternalError(w http.ResponseWriter, err error) {
	Response(w, http.StatusInternalServerError, JSONProps{"message": err})
}
