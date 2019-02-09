package main

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/jean-lopes/data-integration-challenge/companies"
)

func unexpected(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}

func TestData(t *testing.T) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	s, err := companies.CreateService()
	unexpected(t, err)
	unexpected(t, s.Clean())
	loadDatabase(s)
}
