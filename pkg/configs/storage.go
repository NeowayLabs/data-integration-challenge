package configs

import (
	"fmt"
	"github.com/caarlos0/env"
)


type pgConfig struct {
	Host string `env:"PG_HOST" envDefault:"localhost"`
	Port int    `env:"PG_PORT" envDefault:"5432"`
	User string `env:"PG_USER" envDefault:"postgres"`
	Pass string `env:"PG_PASS" envDefault:"postgres"`
}

// ConnectionString builds a connection string based on environment variables
// defaults:
// PG_HOST=localhost
// PG_PORT=5432
// PG_USER=postgres
// PG_PASS=postgres
func ConnectionString() string {
	cfg := pgConfig{}
	env.Parse(&cfg)
	cs := fmt.Sprintf("host=%s port=%d user='%s' password='%s' sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass)
	return cs
}
