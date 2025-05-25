package storage

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/piyusharmap/go-banking/internal/types"
)

type Storage interface {
	RegisterUser(*types.User) (*types.StoredUser, error)
	GetUser(*types.User) (*types.StoredUser, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	return CreateUserTable(s.db)
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

func (s *PostgresStore) RegisterUser(user *types.User) (*types.StoredUser, error) {
	query := `INSERT INTO USERS (contact, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, contact, email, password`

	storedUser := &types.StoredUser{}

	err := s.db.QueryRow(
		query,
		user.Contact,
		user.Email,
		user.Password,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(
		&storedUser.Id,
		&storedUser.Contact,
		&storedUser.Email,
		&storedUser.Password,
	)

	if err != nil {
		return nil, err
	}

	return storedUser, nil
}

func (s *PostgresStore) GetUser(user *types.User) (*types.StoredUser, error) {
	query := `SELECT id, contact, email, password
	FROM USERS
	WHERE contact=$1 AND email=$2`

	storedUser := &types.StoredUser{}

	err := s.db.QueryRow(
		query,
		user.Contact,
		user.Email,
	).Scan(
		&storedUser.Id,
		&storedUser.Contact,
		&storedUser.Email,
		&storedUser.Password,
	)

	if err != nil {
		return nil, err
	}

	return storedUser, nil
}
