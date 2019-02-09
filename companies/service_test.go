package companies

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

type ServiceTestFunc = func(t *testing.T, s Service)

func expectedC(t *testing.T, err *CustomError, specific error) {
	if err == nil {
		t.Fatalf("Unexpected error: %v\n", err)
	}

	var es = []error{}

	if err.Internal != nil {
		es = []error{err.Internal}
	}

	if len(err.Validation) > 0 {
		es = append(err.Validation, es...)
	}

	for _, e := range es {
		if e != nil && e != specific {
			t.Fatalf("Expecting error: %v, got %v", specific, e)
		}
	}
}

func unexpectedC(t *testing.T, err *CustomError) {
	if err != nil {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}

func wrapTestWithService(s Service, sf ServiceTestFunc) func(t *testing.T) {
	return func(t *testing.T) {
		sf(t, s)
	}
}

func testServiceInsert(t *testing.T, s Service) {
	c := TestCompanies[CompanyWithID]
	err := s.Save(&c)
	unexpectedC(t, err)
}

func testServiceInsertEmptyID(t *testing.T, s Service) {
	c := TestCompanies[CompanyNormal]
	err := s.Save(&c)
	unexpectedC(t, err)
}

func testServiceInsertEmptyName(t *testing.T, s Service) {
	c := TestCompanies[CompanyEmptyName]
	err := s.Save(&c)
	expectedC(t, err, ErrEmptyName)
}

func testServiceInsertEmptyZip(t *testing.T, s Service) {
	c := TestCompanies[CompanyEmptyZip]
	err := s.Save(&c)
	expectedC(t, err, ErrEmptyZip)
}

func testServiceInsertZipLetter(t *testing.T, s Service) {
	c := TestCompanies[CompanyInvalidZipDigit]
	err := s.Save(&c)
	expectedC(t, err, ErrInvalidZip)
}

func testServiceInsertZipLength(t *testing.T, s Service) {
	c := TestCompanies[CompanyInvalidZipLength]
	err := s.Save(&c)
	expectedC(t, err, ErrInvalidZip)
}

func testServiceInsertViolateUK(t *testing.T, s Service) {
	c := TestCompanies[CompanyNormal]

	id, err := uuid.NewV4()
	unexpected(t, err)
	c.ID = &id
	c.Name = "UK"
	e := s.Save(&c)
	unexpectedC(t, e)

	id, err = uuid.NewV4()
	unexpected(t, err)
	c.ID = &id
	e = s.Save(&c)
	expectedC(t, e, ErrDuplicate)
}

func testServiceInsertInvalidWebsite(t *testing.T, s Service) {
	c := TestCompanies[CompanyInvalidWebsite]
	err := s.Save(&c)
	expectedC(t, err, ErrInvalidWebsite)
}

func testServiceUpdate(t *testing.T, s Service) {
	c := TestCompanies[CompanyWithID]
	id, err := uuid.NewV4()
	unexpected(t, err)
	c.ID = &id
	c.Name = "Service Update test"
	c.Zip = "12345"

	e := s.Save(&c)
	unexpectedC(t, e)
	ws := "this is website."
	n, err := s.MergeWebsite(&ws, c.Name, c.Zip)
	unexpected(t, err)
	if n != 1 {
		t.Fatal("Update should alter only a single company")
	}
}

func TestService(t *testing.T) {
	s, err := CreateService()
	unexpected(t, err)
	defer s.Close()

	s.Clean()

	m := map[string]ServiceTestFunc{
		"Normal":          testServiceInsert,
		"Empty ID":        testServiceInsertEmptyID,
		"Empty name":      testServiceInsertEmptyName,
		"Empty Zip":       testServiceInsertEmptyZip,
		"Zip Length":      testServiceInsertZipLength,
		"Zip letter":      testServiceInsertZipLetter,
		"Violate UK":      testServiceInsertViolateUK,
		"Invalid website": testServiceInsertInvalidWebsite,
		"Update":          testServiceUpdate,
	}

	for name, f := range m {
		t.Run(name, wrapTestWithService(s, f))
	}
}
