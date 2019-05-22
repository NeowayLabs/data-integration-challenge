package configs

import (
	"fmt"

	"github.com/caarlos0/env"
)

// PgConfig default configuration
type PgConfig struct {
	Host         string `env:"PG_HOST" envDefault:"localhost"`
	Port         int    `env:"PG_PORT" envDefault:"5432"`
	User         string `env:"PG_USER" envDefault:"postgres"`
	Pass         string `env:"PG_PASS" envDefault:"postgres"`
	MaxIdleConns int    `env:"PG_MAX_IDLE_CONNECTIONS" envDefault:"5"`
	MaxOpenConns int    `env:"PG_MAX_OPEN_CONNECTIONS" envDefault:"15"`
}

// ConnectionString builds a connection string based on environment variables
// defaults:
// PG_HOST=localhost
// PG_PORT=5432
// PG_USER=postgres
// PG_PASS=postgres
func ConnectionString(cfg PgConfig) string {
	env.Parse(&cfg)
	cs := fmt.Sprintf("host=%s port=%d user='%s' password='%s' sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass)
	return cs
}
