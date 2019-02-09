package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-http-utils/logger"
	"github.com/jean-lopes/data-integration-challenge/companies"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type handlerFuncWithService = func(w http.ResponseWriter, r *http.Request, s companies.Service)

// JSONProps properties for building a JSON
type JSONProps = map[string]interface{}

func response(w http.ResponseWriter, status int, props JSONProps) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(props)
}

func badRequest(w http.ResponseWriter, errors []error) {
	response(w, http.StatusTeapot, map[string]interface{}{
		"important": "Relax and drink some 🍵 :)",
		"errors":    errors})
}

func internalError(w http.ResponseWriter, err error) {
	log.Printf("Error: %v", err)
	response(w, http.StatusInternalServerError, JSONProps{"message": err})
}

func readJSON(w http.ResponseWriter, r http.Request, c *companies.Company) bool {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		badRequest(w, []error{err})
	}

	return err == nil
}

func updateCompanyHandler(w http.ResponseWriter, r *http.Request, service companies.Service) {
	var c companies.Company

	ok := readJSON(w, *r, &c)
	if !ok {
		return
	}

	rows, err := service.MergeWebsite(c.Website, c.Name, c.Zip)
	if err != nil {
		internalError(w, err)
		return
	}

	if rows == 0 {
		response(w, http.StatusNoContent, JSONProps{})
	} else {
		response(w, http.StatusOK, JSONProps{"affected rows": rows})
	}
}

func createCompanyHandler(w http.ResponseWriter, r *http.Request, service companies.Service) {
	var c companies.Company

	ok := readJSON(w, *r, &c)
	if !ok {
		return
	}

	customError := service.Save(&c)
	if customError != nil {
		if customError.Internal != nil {
			internalError(w, customError.Internal)
		} else {
			badRequest(w, customError.Validation)
		}
		return
	}

	response(w, http.StatusCreated, JSONProps{"id": c.ID})
}

func wrapHandler(s companies.Service, m string, h handlerFuncWithService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Compare(strings.ToUpper(r.Method), strings.ToUpper(m)) != 0 {
			response(w, http.StatusMethodNotAllowed, JSONProps{})
		} else {
			h(w, r, s)
		}
	}
}

// NewMux creates the handlers
func NewMux(s companies.Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/companies/create", wrapHandler(s, "POST", createCompanyHandler))
	mux.HandleFunc("/companies/update", wrapHandler(s, "PUT", updateCompanyHandler))
	return mux
}

func main() {
	service, err := companies.CreateService()
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()

	err = service.Clean()
	if err != nil {
		log.Fatal(err)
	}

	loadDatabase(service)

	mux := NewMux(service)

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", logger.Handler(mux, os.Stdout, logger.TinyLoggerType)))
}
