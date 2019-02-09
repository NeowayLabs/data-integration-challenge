package companies

import (
	"errors"
	"log"

	uuid "github.com/satori/go.uuid"
)

// CustomError separates internal errors from validation errors
type CustomError struct {
	Validation []error
	Internal   error
}

// Service for companies manipulation
type Service interface {
	Save(c *Company) *CustomError
	MergeWebsite(ws *string, name string, zip string) (int64, error)
	Clean() error
	Close()
}

// Companies struct representing the service interface
type Companies struct {
	r *repository
}

// CreateService return a new instance of the Service interface
func CreateService() (Service, error) {
	r, err := createRepository()
	var s Service
	s = Companies{r: r}
	return s, err
}

// Clean the all the data in database
func (cs Companies) Clean() error {
	return cs.r.clean()
}

// Close the service
func (cs Companies) Close() {
	cs.r.close()
}

// Save inserts a company in the database
func (cs Companies) Save(c *Company) *CustomError {
	customError := CustomError{}

	if c == nil {
		customError.Internal = errors.New("Cannot create nil Company")
		return &customError
	}

	if !c.HasID() {
		id, err := generateID()
		if err != nil {
			customError.Internal = err
			return &customError
		}

		c.ID = &id
	}

	customError.Validation = c.Validate()

	if len(customError.Validation) == 0 {
		err := cs.r.insert(*c)
		if err == ErrDuplicate {
			customError.Validation = appendError(customError.Validation, err)
			return &customError
		}

		if err != nil {
			customError.Internal = cs.r.insert(*c)
		}
	}

	if customError.Internal == nil && len(customError.Validation) == 0 {
		return nil
	}

	return &customError
}

// MergeWebsite updates a company's website, return the number of rows updated
func (cs Companies) MergeWebsite(ws *string, name string, zip string) (int64, error) {
	return cs.r.updateWebsite(ws, name, zip)
}

func generateID() (uuid.UUID, error) {
	id, err := uuid.NewV4()

	if err != nil {
		log.Printf("Failed to generate UUID v4: %s\n", err)
	}

	return id, err
}
