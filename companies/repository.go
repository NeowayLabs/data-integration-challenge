package companies

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/lib/pq"
)

type pgConfig struct {
	Host string `env:"PG_HOST" envDefault:"localhost"`
	Port int    `env:"PG_PORT" envDefault:"5432"`
	User string `env:"PG_USER" envDefault:"postgres"`
	Pass string `env:"PG_PASS" envDefault:"postgres"`
}

type repository struct {
	db         *sql.DB
	insertStmt *sql.Stmt
	updateStmt *sql.Stmt
}

// ErrDuplicate violation of unique key
var ErrDuplicate = errors.New("Record already exists")

// ConnectionString builds a connection string based on environment variables
// defaults:
// PG_HOST=localhost
// PG_PORT=5432
// PG_USER=postgres
// PG_PASS=postgres
func connectionString() string {
	cfg := pgConfig{}
	env.Parse(&cfg)
	cs := fmt.Sprintf("host=%s port=%d user='%s' password='%s' sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass)
	return cs
}

// Setup database connection and prepared statements
func createRepository() (*repository, error) {
	db, err := sql.Open("postgres", connectionString())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(15)

	insert, err := db.Prepare(`INSERT INTO companies(id, name, zip, website) values($1, UPPER($2), $3, LOWER($4))`)
	if err != nil {
		return nil, err
	}

	update, err := db.Prepare(`UPDATE companies SET website = LOWER($1) WHERE name = UPPER($2) and zip = $3`)
	if err != nil {
		return nil, err
	}

	r := repository{db: db, insertStmt: insert, updateStmt: update}

	return &r, nil
}

func (r repository) close() {
	r.insertStmt.Close()
	r.updateStmt.Close()
	r.db.Close()
}

// Insert a company, unique key is not checked before inserting
// ID, Name, Zip are required fields
// Required fields are not checked
func (r repository) insert(c Company) error {
	_, err := r.insertStmt.Exec(c.ID, c.Name, c.Zip, c.Website)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == pq.ErrorCode("23505") {
				return ErrDuplicate
			}
		}
	}

	return err
}

// UpdateWebsite updates a company website.
// returns the number of rows affected by the update
func (r repository) updateWebsite(website *string, name string, zip string) (int64, error) {
	res, err := r.updateStmt.Exec(website, name, zip)
	if err != nil {
		return 0, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}

func (r repository) clean() error {
	_, err := r.db.Exec("TRUNCATE companies")
	return err
}
