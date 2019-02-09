package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jean-lopes/data-integration-challenge/companies"
)

type testData struct {
	method  string
	path    string
	status  int
	reqBody string
	resBody string
}

var (
	allData = map[string]testData{
		"CREATE - Normal": {
			method:  "POST",
			path:    "/companies/create",
			status:  http.StatusCreated,
			reqBody: `{"id":"60992c98-7228-424e-8727-0eafa9351054","name":"b","zip":"00000"}`,
			resBody: `{"id":"60992c98-7228-424e-8727-0eafa9351054"}`},
		"CREATE - Bad Request": {
			method:  "POST",
			path:    "/companies/create",
			status:  http.StatusTeapot,
			reqBody: `{"name": "","zip":"01234"}`,
			resBody: ``},
		"CREATE - Method not allowed": {
			method:  "GET",
			path:    "/companies/create",
			status:  http.StatusMethodNotAllowed,
			reqBody: `{}`,
			resBody: `{}`},
		"UPDATE - OK": {
			method:  "PUT",
			path:    "/companies/update",
			status:  http.StatusOK,
			reqBody: `{"website":"https://duckduckgo.com","name":"12345","zip":"12345"}`,
			resBody: `{"affected rows":1}`},
		"UPDATE - No content": {
			method:  "PUT",
			path:    "/companies/update",
			status:  http.StatusNoContent,
			reqBody: `{"website":"https://duckduckgo.com","name":"non existent","zip":"12345"}`,
			resBody: `{}`},
		"UPDATE - Bad request": {
			method:  "PUT",
			path:    "/companies/update",
			status:  http.StatusTeapot,
			reqBody: ``,
			resBody: ``},
		"UPDATE - Method not allowed": {
			method:  "GET",
			path:    "/companies/update",
			status:  http.StatusMethodNotAllowed,
			reqBody: `{}`,
			resBody: `{}`},
	}
)

func newRequest(t *testing.T, data testData) *http.Request {
	return httptest.NewRequest(data.method, data.path, strings.NewReader(data.reqBody))
}

func TestAPI(t *testing.T) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	service, err := companies.CreateService()
	if err != nil {
		t.Fatal(err)
	}
	defer service.Close()

	service.Clean()

	mux := NewMux(service)

	service.Clean()
	if err != nil {
		t.Fatal(err)
	}

	e := service.Save(&companies.Company{Name: "12345", Zip: "12345"})
	if e != nil {
		t.Fatalf("Unexpected error: %v\n", e)
	}

	for name, data := range allData {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRequest(data.method, data.path, strings.NewReader(data.reqBody))
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			res := w.Result()

			if res.StatusCode != data.status {
				t.Fatalf("Expected status code %v, got %v - %v", data.status, res.StatusCode, res)
			}

			if data.resBody == "" {
				return
			}

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Unexpected error: %v\n", err)
			}
			res.Body.Close()

			body := strings.TrimSpace(string(bs))

			if strings.Compare(data.resBody, body) != 0 {
				t.Fatalf("Expected %v, got %v", data.resBody, body)
			}
		})
	}
}
