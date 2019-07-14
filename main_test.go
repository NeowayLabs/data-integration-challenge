package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/coscms/xorm"
	"github.com/ruiblaese/data-integration-challenge/db"
	"github.com/ruiblaese/data-integration-challenge/models"
	"github.com/ruiblaese/data-integration-challenge/routes"
	"github.com/ruiblaese/data-integration-challenge/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

var xormEngine *xorm.Engine

func performRequest(r http.Handler, method, path, json string,
	t *testing.T) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	var req *http.Request

	if json != "" {
		req, _ = http.NewRequest(method, path, bytes.NewReader([]byte(json)))

	} else {
		req, _ = http.NewRequest(method, path, nil)
	}
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	return w
}

func performRequestFile(r http.Handler, method, path, filePath string,
	t *testing.T) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	var req *http.Request

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, file)
	writer.Close()

	req, _ = http.NewRequest(method, path, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	r.ServeHTTP(w, req)

	return w
}

func TestStartDatabase(t *testing.T) {

	fileDB := "./challenge.db"
	if _, err := os.Stat(fileDB); !os.IsNotExist(err) {
		os.Remove(fileDB)
	}

	xorm, err := db.StartDatabase()
	assert.NoError(t, err)

	xormEngine = xorm
}

func TestProcessesFirstData(t *testing.T) {
	err := services.ProcessesFirstData(xormEngine)
	assert.NoError(t, err)
}

func TestCompaniesWithServer(t *testing.T) {

	ginRouter := gin.Default()
	router := routes.StartRouter(ginRouter, xormEngine)

	t.Run("GET_All", func(t *testing.T) {
		w := performRequest(router, http.MethodGet, "/api/v1/companies", "", t)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("POST_Insert", func(t *testing.T) {
		w := performRequest(
			router,
			http.MethodPost,
			"/api/v1/companies",
			`{"name":"company test","zip":"12345"}`,
			t)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("POST_Upload", func(t *testing.T) {

		folder, err := services.ProcessesChallengeFiles()
		assert.NoError(t, err)

		w := performRequestFile(
			router,
			http.MethodPost,
			"/api/v1/companies/upload",
			folder+"q2_clientData.csv",
			t)
		assert.Equal(t, http.StatusOK, w.Code)
		os.RemoveAll(folder)
	})

	t.Run("GET_Name_Parcial", func(t *testing.T) {
		w := performRequest(router, http.MethodGet, "/api/v1/companies?name=company", "", t)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET_Name_Parcial_AND_zipcode", func(t *testing.T) {
		w := performRequest(router, http.MethodGet, "/api/v1/companies?name=company&zipcode=12345", "", t)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("DELETE", func(t *testing.T) {
		w := performRequest(router, http.MethodGet, "/api/v1/companies?name=company&zipcode=12345", "", t)
		assert.Equal(t, http.StatusOK, w.Code)

		bodyBytes, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)

		companies := []models.Company{}
		err = json.Unmarshal(bodyBytes, &companies)
		assert.NoError(t, err)

		w = performRequest(router, http.MethodDelete, "/api/v1/companies/"+strconv.Itoa(companies[0].ID), "", t)
		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("get by id, not found", func(t *testing.T) {

		w := performRequest(router, http.MethodGet, "/api/v1/companies/999999", "", t)
		assert.Equal(t, http.StatusNotFound, w.Code)

		resBytes, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)
		assert.Empty(t, resBytes)
	})

}
