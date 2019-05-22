package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jean-lopes/data-integration-challenge/pkg/httphelpers"
	"github.com/jean-lopes/data-integration-challenge/pkg/models"
	"github.com/jean-lopes/data-integration-challenge/pkg/storage"
)

// MergeCompanyWebsiteHandlerBody nice long name k
type MergeCompanyWebsiteHandlerBody struct {
	Name    string  `json:"name"`
	Zip     string  `json:"zip"`
	Website *string `json:"website"`
}

// MergeCompanyWebsiteHandler handler for updating companies .-.
func MergeCompanyWebsiteHandler(store storage.Company) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		args := MergeCompanyWebsiteHandlerBody{}

		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			httphelpers.BadRequest(w, []error{err})
			return
		}

		c := models.Company{Name: args.Name, Zip: args.Zip, Website: args.Website}
		errs := c.Validate()
		if errs != nil {
			httphelpers.BadRequest(w, errs)
			return
		}

		err = store.UpdateWebsite(args.Website, args.Name, args.Zip)
		if err != nil {
			httphelpers.InternalError(w, err)
			return
		}

		httphelpers.Response(w, http.StatusOK, httphelpers.JSONProps{})
	}
}
