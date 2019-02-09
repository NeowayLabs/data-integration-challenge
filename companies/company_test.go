package companies

import (
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
)

const (
	// CompanyNormal key for TestCompanies map
	CompanyNormal = "normal"
	// CompanyWithID key for TestCompanies map
	CompanyWithID = "company-with-id"
	// CompanyEmptyName key for TestCompanies map
	CompanyEmptyName = "company-with-empty-name"
	// CompanyEmptyZip key for TestCompanies map
	CompanyEmptyZip = "company-with-empty-zip"
	// CompanyWithInvalidZipDigit key for TestCompanies map
	CompanyInvalidZipDigit = "company-with-invalid-zip-digit"
	// CompanyWithInvalidZeipLength key for TestCompanies map
	CompanyInvalidZipLength = "company-with-invalid-zip-length"
	// CompanyInvalidWebsite key for TestCompanies map
	CompanyInvalidWebsite = "company-with-invalid-website"
	// CompanyAllWrong key for TestCompanies map
	CompanyAllWrong = "company-all-wrong"
)

var (
	invalidWebsiteStr = "Not a valid website :)"
	// CompanyID used for tests
	CompanyID = uuid.Must(uuid.FromString("60992c98-7228-424e-8727-0eafa9351054"))
	// TestCompanies Map of companies for use in tests
	TestCompanies = map[string]Company{
		CompanyNormal: Company{
			ID:      nil,
			Name:    "a",
			Zip:     "00000",
			Website: nil},
		CompanyWithID: Company{
			ID:      &CompanyID,
			Name:    "b",
			Zip:     "00000",
			Website: nil},
		CompanyEmptyName: Company{
			ID:      nil,
			Name:    "",
			Zip:     "00000",
			Website: nil},
		CompanyEmptyZip: Company{
			ID:      nil,
			Name:    "a",
			Zip:     "",
			Website: nil},
		CompanyInvalidZipDigit: Company{
			ID:      nil,
			Name:    "a",
			Zip:     "a",
			Website: nil},
		CompanyInvalidZipLength: Company{
			ID:      nil,
			Name:    "a",
			Zip:     "0",
			Website: nil},
		CompanyInvalidWebsite: Company{
			ID:      nil,
			Name:    "a",
			Zip:     "00000",
			Website: &invalidWebsiteStr},
		CompanyAllWrong: Company{
			ID:      nil,
			Name:    "",
			Zip:     "",
			Website: &invalidWebsiteStr},
	}
)

func TestCompanyHasID(t *testing.T) {
	noID := TestCompanies[CompanyNormal]
	withID := TestCompanies[CompanyWithID]

	if noID.HasID() {
		t.Fatalf("Expecting HasID() to be false, company: %v", noID)
	}

	if !withID.HasID() {
		t.Fatalf("Expecting HasID() to be true, company: %v", withID)
	}
}

func TestCompanyIsEmpty(t *testing.T) {
	c := TestCompanies[CompanyNormal]

	if c.IsEmpty() {
		t.Fatalf("%v.isEmpty() should return false", c)
	}

	if !Nil.IsEmpty() {
		t.Fatalf("%v.IsEmpty() should be true", Nil)
	}
}

func TestCompanyEqual(t *testing.T) {
	c := TestCompanies[CompanyNormal]
	i := TestCompanies[CompanyWithID]

	if c.Equal(Nil) {
		t.Fatalf("%v should be different than %v", c, Nil)
	}

	if Nil.Equal(c) {
		t.Fatalf("%v should be different than %v", Nil, c)
	}

	if c.Equal(i) {
		t.Fatalf("%v should be different than %v", c, i)

	}

	if i.Equal(c) {
		t.Fatalf("%v should be different than %v", i, c)
	}

	if !i.Equal(i) {
		t.Fatalf("%v should be equal to %v", i, i)
	}
}

func TestCompanyValidate(t *testing.T) {
	normal := TestCompanies[CompanyNormal]
	emptyName := TestCompanies[CompanyEmptyName]
	emptyZip := TestCompanies[CompanyEmptyZip]
	invalidZipDigit := TestCompanies[CompanyInvalidZipDigit]
	invalidZipLength := TestCompanies[CompanyInvalidZipLength]
	invalidWebsite := TestCompanies[CompanyInvalidWebsite]
	allWrong := TestCompanies[CompanyAllWrong]

	m := map[string]struct {
		c    Company
		errs []error
	}{
		"Normal":                        {c: normal, errs: nil},
		"Empty name":                    {c: emptyName, errs: []error{ErrEmptyName}},
		"Empty zip":                     {c: emptyZip, errs: []error{ErrEmptyZip}},
		"Empty invalid zip (non digit)": {c: invalidZipDigit, errs: []error{ErrInvalidZip}},
		"Empty invalid zip (length)":    {c: invalidZipLength, errs: []error{ErrInvalidZip}},
		"Invalid website":               {c: invalidWebsite, errs: []error{ErrInvalidWebsite}},
		"Every error":                   {c: allWrong, errs: []error{ErrEmptyName, ErrEmptyZip, ErrInvalidWebsite}},
	}

	f := func(c Company, errs *[]error) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()
			if reflect.DeepEqual(c.Validate(), errs) {
				t.Fatalf("%v.Validate() should return %v", c, errs)
			}
		}
	}

	for key, value := range m {
		v := value
		t.Run(key, f(v.c, &v.errs))
	}
}
