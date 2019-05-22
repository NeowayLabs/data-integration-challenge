package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jean-lopes/data-integration-challenge/pkg/httphelpers"
	"github.com/jean-lopes/data-integration-challenge/pkg/models"
	"github.com/jean-lopes/data-integration-challenge/pkg/storage"
	uuid "github.com/satori/go.uuid"
)

// CreateCompanyHandler handler for creating companies .-.
func CreateCompanyHandler(store storage.Company) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		c := models.Company{}

		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			httphelpers.BadRequest(w, []error{err})
			return
		}

		errs := c.Validate()
		if errs != nil {
			httphelpers.BadRequest(w, errs)
			return
		}

		exists, err := store.Exists(c.Name, c.Zip)
		if err != nil {
			httphelpers.InternalError(w, err)
		}

		if exists {
			httphelpers.BadRequest(w, []error{models.ErrCompanyAlreadyExists})
			return
		}

		id := uuid.NewV4()
		c.ID = &id
		err = store.Save(c)
		if err != nil {
			httphelpers.InternalError(w, err)
		}

		httphelpers.Response(w, 200, httphelpers.JSONProps{"id": c.ID})
	}
}
