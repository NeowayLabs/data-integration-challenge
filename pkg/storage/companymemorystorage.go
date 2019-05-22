package storage

import (
	"strings"

	"github.com/jean-lopes/data-integration-challenge/pkg/models"
	uuid "github.com/satori/go.uuid"
)

// InMemory Memory storage implementation
type InMemory struct {
	es map[uuid.UUID]models.Company
	uk map[string]uuid.UUID
}

// OpenInMemoryStorage a PostgreSQL storage
func OpenInMemoryStorage() (Company, error) {
	c := InMemory{
		make(map[uuid.UUID]models.Company),
		make(map[string]uuid.UUID),
	}

	return c, nil
}

// "hash"
func hash(c models.Company) string {
	return strings.ToUpper(c.Name) + c.Zip
}

// Close the storage
func (im InMemory) Close() {
	im.Clean()
}

// Exists checks for existance of a company
func (im InMemory) Exists(name string, zip string) (bool, error) {
	_, exists := im.uk[hash(models.Company{Name: name, Zip: zip})]
	return exists, nil
}

// Save a company, unique key is not checked before inserting
func (im InMemory) Save(c models.Company) error {
	im.es[*c.ID] = c
	im.uk[hash(c)] = *c.ID
	return nil
}

// SaveAll companies from the channel
func (im InMemory) SaveAll(companies <-chan models.Company) error {
	if companies != nil {
		for c := range companies {
			im.Save(c)
		}
	}

	return nil
}

// UpdateWebsite No validations made
func (im InMemory) UpdateWebsite(website *string, name string, zip string) error {
	id, exists := im.uk[hash(models.Company{Name: name, Zip: zip})]
	if exists {
		c := im.es[id]
		c.Website = website
	}
	return nil
}

// Clean all companies from the storage
func (im InMemory) Clean() error {
	im.es = make(map[uuid.UUID]models.Company)
	im.uk = make(map[string]uuid.UUID)
	return nil
}
