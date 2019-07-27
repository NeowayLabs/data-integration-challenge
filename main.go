package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/company/", GetCompanies).Methods("GET")
    router.HandleFunc("/company/{id}", GetCompany).Methods("GET")
    router.HandleFunc("/company/{id}", CreateCompany).Methods("POST")
    router.HandleFunc("/company/{id}", DeleteCompany).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", router))
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {}
func GetCompany(w http.ResponseWriter, r *http.Request) {}
func CreateCompany(w http.ResponseWriter, r *http.Request) {}
func DeleteCompany(w http.ResponseWriter, r *http.Request) {}
