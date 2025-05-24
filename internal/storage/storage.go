package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/piyusharmap/go-banking/internal/types"
)

type Storage interface {
	RegisterUser(*types.User) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) RegisterUser(*types.User) error {
	return nil
}
