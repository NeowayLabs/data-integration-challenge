package csv

import (
	"log"

	"github.com/jean-lopes/data-integration-challenge/pkg/models"
)

// ToCompanies converts the channel data to the company model
func ToCompanies(records <-chan []string, chanSize int) <-chan models.Company {
	ch := make(chan models.Company, chanSize)

	go func() {
		defer close(ch)

		for rec := range records {
			recSize := len(rec)

			// TODO verificar conteudo do slice (não só o tamanho)
			if recSize >= 2 && recSize <= 3 {
				name := rec[0]
				zip := rec[1]
				var ws *string

				if len(rec) == 3 {
					ws = &rec[2]
				}

				c := models.Company{Name: name, Zip: zip, Website: ws}

				ch <- c
			} else {
				log.Printf("Error: invalid CSV line: '%v'", rec) // TODO adicionar linha do CSV no log
			}
		}
	}()

	return ch
}
