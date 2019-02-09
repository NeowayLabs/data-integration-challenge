package companies

import (
	"log"
	"testing"

	uuid "github.com/satori/go.uuid"
)

type repositoryTestFunc = func(t *testing.T, r repository)

func unexpected(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}

func expected(t *testing.T, err error) {
	if err == nil {
		log.Fatal("Expected error, and got none")
	}
}

func wrapTestWithRepository(r repository, rf repositoryTestFunc) func(t *testing.T) {
	return func(t *testing.T) {
		rf(t, r)
	}
}

func testInsert(t *testing.T, r repository) {
	c := TestCompanies[CompanyWithID]
	err := r.insert(c)
	unexpected(t, err)
}

func testInsertEmptyID(t *testing.T, r repository) {
	c := TestCompanies[CompanyNormal]
	err := r.insert(c)
	expected(t, err)
}

func testInsertEmptyName(t *testing.T, r repository) {
	c := TestCompanies[CompanyEmptyName]
	err := r.insert(c)
	expected(t, err)
}

func testInsertEmptyZip(t *testing.T, r repository) {
	c := TestCompanies[CompanyEmptyZip]
	err := r.insert(c)
	expected(t, err)
}

func testInsertZipLength(t *testing.T, r repository) {
	c := TestCompanies[CompanyInvalidZipLength]
	err := r.insert(c)
	expected(t, err)
}

func testInsertViolateUK(t *testing.T, r repository) {
	c := TestCompanies[CompanyNormal]

	id, err := uuid.NewV4()
	c.ID = &id
	c.Name = "UK"
	err = r.insert(c)
	unexpected(t, err)

	id, err = uuid.NewV4()
	c.ID = &id
	err = r.insert(c)
	expected(t, err)
}

func testUpdate(t *testing.T, r repository) {
	id, err := uuid.NewV4()
	unexpected(t, err)
	c := Company{ID: &id, Name: "update", Zip: "01234", Website: nil}
	err = r.insert(c)
	unexpected(t, err)
	ws := "a"
	n, err := r.updateWebsite(&ws, "update", "01234")
	unexpected(t, err)
	if n != 1 {
		t.Fatal("Update should alter only a single company")
	}
}

func TestRepository(t *testing.T) {
	r, err := createRepository()
	unexpected(t, err)
	defer r.close()

	r.clean()

	m := map[string]repositoryTestFunc{
		"Normal":     testInsert,
		"Empty ID":   testInsertEmptyID,
		"Empty name": testInsertEmptyName,
		"Empty Zip":  testInsertEmptyZip,
		"ZipLength":  testInsertZipLength,
		"ViolateUK":  testInsertViolateUK,
		"update":     testUpdate,
	}

	for name, f := range m {
		t.Run(name, wrapTestWithRepository(*r, f))
	}
}
