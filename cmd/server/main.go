package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-http-utils/logger"
	"github.com/jean-lopes/data-integration-challenge/pkg/configs"
	"github.com/jean-lopes/data-integration-challenge/pkg/handlers"
	"github.com/jean-lopes/data-integration-challenge/pkg/httphelpers"
	"github.com/jean-lopes/data-integration-challenge/pkg/storage"
)

func wrapHandler(m string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Compare(strings.ToUpper(r.Method), strings.ToUpper(m)) != 0 {
			httphelpers.Response(w, http.StatusMethodNotAllowed, httphelpers.JSONProps{})
		} else {
			h(w, r)
		}
	}
}

func main() {
	store, err := storage.OpenPostgreSQLStorage(configs.PgConfig{})
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	createCompanyHandler := handlers.CreateCompanyHandler(store)
	mergeCompanyWebsiteHandler := handlers.MergeCompanyWebsiteHandler(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/companies", wrapHandler("POST", createCompanyHandler))
	mux.HandleFunc("/merge-company-website", wrapHandler("POST", mergeCompanyWebsiteHandler))

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", logger.Handler(mux, os.Stdout, logger.TinyLoggerType)))
}
