package storage

import (
	"github.com/jean-lopes/data-integration-challenge/pkg/csv"
	"github.com/jean-lopes/data-integration-challenge/pkg/models"
	uuid "github.com/satori/go.uuid"
)

// LoadFromCSV Load data from CSV file without headers into the storage
func LoadFromCSV(store Company, path string) error {
	channelSize := 10

	records := csv.ReadAll(path, ';', channelSize)

	noIDs := csv.ToCompanies(records, channelSize)

	companies := models.MapCompanies(noIDs, channelSize, func(c models.Company) models.Company {
		id := uuid.NewV4()
		c.ID = &id
		return c
	})

	return store.SaveAll(companies)
}
