package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nolint

	"github.com/desmos-labs/plutus/types"
)

// Database represents the source inside which the data will be stored
type Database struct {
	sql *sqlx.DB
}

// NewDatabase allows to create a new Database instance
func NewDatabase(cfg *types.DatabaseConfig) (*Database, error) {
	postgresDb, err := sqlx.Connect("postgres", cfg.URI)
	if err != nil {
		return nil, err
	}

	return &Database{sql: postgresDb}, nil
}
