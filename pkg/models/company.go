package models

import (
	"errors"
	"net/url"
	"regexp"

	"github.com/jean-lopes/data-integration-challenge/pkg/util"
	uuid "github.com/satori/go.uuid"
)

// Company entity
type Company struct {
	ID      *uuid.UUID `json:"id,omitempty"`
	Name    string     `json:"name"`
	Zip     string     `json:"zip"`
	Website *string    `json:"website,omitempty"`
}

// CompanyMapper generic company mapper
type CompanyMapper = func(Company) Company

var (
	// ErrEmptyName empty name error
	ErrEmptyName = errors.New("Empty company name")
	// ErrEmptyZip empty zip error
	ErrEmptyZip = errors.New("Empty company zip code")
	// ErrInvalidZip Invalid zip error
	ErrInvalidZip = errors.New("Company zip code must have exacly 5 (five) digits")
	// ErrInvalidWebsite malformed URI
	ErrInvalidWebsite = errors.New("Invalid website")
	// ErrCompanyAlreadyExists Company already exists on storage
	ErrCompanyAlreadyExists = errors.New("Company already exists")
	zipRE                   = regexp.MustCompile("[0-9]{5}")
)

// Validate business constraints
func (company Company) Validate() []error {
	nameError := validateName(company.Name)
	zipError := validateZip(company.Zip)
	websiteError := validateWebsite(company.Website)

	errors := make([]error, 0)
	errors = util.AppendError(errors, nameError)
	errors = util.AppendError(errors, zipError)
	errors = util.AppendError(errors, websiteError)

	if len(errors) == 0 {
		return nil
	}

	return errors
}

// HasID verify if the company has ID
func (company Company) HasID() bool {
	return company.ID != nil && !uuid.Equal(*company.ID, uuid.Nil)
}

// IsEmpty checks if an instance of company is empty
func (company Company) IsEmpty() bool {
	return company.ID == nil &&
		company.Name == "" &&
		company.Zip == "" &&
		company.Website == nil
}

// Equal Checks if two companies instances have the same ID
func (company Company) Equal(other Company) bool {
	return company.ID != nil && other.ID != nil && uuid.Equal(*company.ID, *other.ID)
}

func validateName(name string) error {
	if util.IsBlank(name) {
		return ErrEmptyName
	}

	return nil
}

func validateZip(zip string) error {
	if util.IsBlank(zip) {
		return ErrEmptyZip
	}

	if !zipRE.MatchString(zip) {
		return ErrInvalidZip
	}

	return nil
}

func validateWebsite(website *string) error {
	if website != nil && !util.IsBlank(*website) {
		_, err := url.ParseRequestURI(*website)
		if err != nil {
			return ErrInvalidWebsite
		}
	}

	return nil
}

// MapCompanies company stream
func MapCompanies(companies <-chan Company, channelSize int, fn CompanyMapper) chan Company {
	ch := make(chan Company, channelSize)

	go func() {
		defer close(ch)

		for c := range companies {
			c = fn(c)

			ch <- c
		}
	}()

	return ch
}
