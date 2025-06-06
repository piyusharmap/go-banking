package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/piyusharmap/go-banking/internal/types"
)

type Storage interface {
	// customer methods
	RegisterCustomer(*types.Customer) (*types.CustomerResponse, error)
	GetCustomer(*types.Customer) (*types.CustomerModel, error)
	GetCustomerByID(int) (*types.CustomerResponse, error)
	UpdateCustomer(int, *types.UpdateCustomerRequest) (*types.CustomerResponse, error)
	DeleteCustomer(int) (*types.CustomerResponse, error)

	// account methods
	RegisterAccount(*types.Account) (*types.AccountResponse, error)
	GetAccountByID(int, int) (*types.AccountResponse, error)
	UpdateAccount(int, int, *types.UpdateAccountRequest) (*types.AccountResponse, error)
	RemoveAccount(int, int) (*types.AccountResponse, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	return s.db.Ping()
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
