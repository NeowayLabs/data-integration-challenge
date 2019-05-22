package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PG

	"github.com/jean-lopes/data-integration-challenge/pkg/configs"
	"github.com/jean-lopes/data-integration-challenge/pkg/models"
	"github.com/jean-lopes/data-integration-challenge/pkg/util"
)

// Pg PostgreSQL storage implementation
type Pg struct {
	db         *sql.DB
	insertStmt *sql.Stmt
	updateStmt *sql.Stmt
	deleteStmt *sql.Stmt
	existsStmt *sql.Stmt
}

// OpenPostgreSQLStorage a PostgreSQL storage
func OpenPostgreSQLStorage(pgConfig configs.PgConfig) (Company, error) {
	cs := configs.ConnectionString(pgConfig)
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(pgConfig.MaxIdleConns)
	db.SetMaxOpenConns(pgConfig.MaxOpenConns)

	insert, err := db.Prepare(`INSERT INTO companies(id, name, zip, website) values($1, UPPER($2), $3, LOWER($4))`)
	if err != nil {
		return nil, err
	}

	update, err := db.Prepare(`UPDATE companies SET website = LOWER($1) WHERE name = UPPER($2) and zip = $3`)
	if err != nil {
		return nil, err
	}

	exists, err := db.Prepare(`SELECT count(id) FROM companies WHERE name = UPPER($1) AND zip = $2`)
	if err != nil {
		return nil, err
	}

	pg := Pg{db: db, insertStmt: insert, updateStmt: update, existsStmt: exists}

	return pg, nil
}

// Close the storage
func (pg Pg) Close() {
	pg.insertStmt.Close()
	pg.updateStmt.Close()
	pg.db.Close()
}

// Exists checks for existance of a company
func (pg Pg) Exists(name string, zip string) (bool, error) {
	row := pg.existsStmt.QueryRow(name, zip)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Save a company, unique key is not checked before inserting
func (pg Pg) Save(c models.Company) error {
	_, err := pg.insertStmt.Exec(c.ID, c.Name, c.Zip, c.Website)
	return err
}

// SaveAll companies from the channel
func (pg Pg) SaveAll(companies <-chan models.Company) error {
	if companies == nil {
		return nil
	}

	// TODO using COPY is a lot faster than batch inserts
	tx, err := pg.db.Begin()
	if err != nil {
		return err
	}

	insertTx := tx.Stmt(pg.insertStmt)

	// TODO a single commit for the whole block might not be ideal, check expected input volume
	for c := range companies {
		errs := c.Validate()
		if errs != nil {
			e := util.MergeErrors(errs)
			log.Printf("Ignoring %v. Failed validations with: %s", c, e)
			continue
		}

		_, err := insertTx.Exec(c.ID, c.Name, c.Zip, c.Website)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// UpdateWebsite No validations made
func (pg Pg) UpdateWebsite(website *string, name string, zip string) error {
	_, err := pg.updateStmt.Exec(website, name, zip)
	return err
}

// Clean all companies from the storage
func (pg Pg) Clean() error {
	_, err := pg.db.Exec("TRUNCATE companies")
	return err
}
