package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/jean-lopes/data-integration-challenge/companies"
)

type converter = func([]string) companies.Company

func readCSV(path string) chan companies.Company {
	cs := make(chan companies.Company, 10)

	go func() {
		defer close(cs)
		fr, err := os.Open(path)
		if err != nil {
			close(cs)
			panic(err)
		}
		defer fr.Close()

		r := csv.NewReader(bufio.NewReader(fr))
		r.Comma = ';'

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				close(cs)
				log.Fatal(err)
			}

			name := record[0]
			zip := record[1]
			var ws *string
			if len(record) == 3 {
				ws = &record[2]
			}

			c := companies.Company{Name: name, Zip: zip, Website: ws}

			cs <- c
		}
	}()

	return cs
}

func loadDatabase(service companies.Service) {
	cs := readCSV("q1_catalog.csv")

	var line = 0
	for c := range cs {
		line++
		err := service.Save(&c)
		if err != nil {
			log.Printf("Ignoring line %d from %q. Error: %v", line, "q1_catalog.csv", err)
		}
	}
}
