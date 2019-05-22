package storage

import "github.com/jean-lopes/data-integration-challenge/pkg/models"

// Company storage interface
type Company interface {
	// Exists checks for existance of a company
	Exists(name string, zip string) (bool, error)
	// Save No validations made
	Save(c models.Company) error
	// SaveAll No validations made
	SaveAll(companies <-chan models.Company) error
	// UpdateWebsite No validations made
	UpdateWebsite(website *string, name string, zip string) error
	// Close the storage
	Close()
	// Clean the storage
	Clean() error
}
