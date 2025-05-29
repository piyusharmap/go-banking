package storage

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/piyusharmap/go-banking/internal/types"
)

type Storage interface {
	RegisterUser(*types.User) (*types.UserResponse, error)
	GetUser(*types.User) (*types.UserModel, error)
	GetUserByID(int) (*types.UserResponse, error)
	UpdateUser(int, *types.User) (*types.UserResponse, error)
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

func (s *PostgresStore) RegisterUser(user *types.User) (*types.UserResponse, error) {
	query := `INSERT INTO users (contact, email, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, contact, email`

	response := &types.UserResponse{}

	err := s.db.QueryRow(
		query,
		user.Contact,
		user.Email,
		user.Password,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(
		&response.ID,
		&response.Contact,
		&response.Email,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostgresStore) GetUser(user *types.User) (*types.UserModel, error) {
	query := `SELECT id, contact, email, password_hash
	FROM users
	WHERE contact=$1 AND email=$2`

	storedUser := &types.UserModel{}

	err := s.db.QueryRow(
		query,
		user.Contact,
		user.Email,
	).Scan(
		&storedUser.ID,
		&storedUser.Contact,
		&storedUser.Email,
		&storedUser.Password,
	)

	if err != nil {
		return nil, err
	}

	return storedUser, nil
}

func (s *PostgresStore) GetUserByID(id int) (*types.UserResponse, error) {
	query := `SELECT id, contact, email
	FROM users
	WHERE id=$1`

	response := &types.UserResponse{}

	err := s.db.QueryRow(
		query,
		id,
	).Scan(
		&response.ID,
		&response.Contact,
		&response.Email,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostgresStore) UpdateUser(id int, user *types.User) (*types.UserResponse, error) {
	query := `UPDATE users
	SET contact=$1, email=$2, updated_at=$3
	WHERE id=$4
	RETURNING id, contact, email`

	response := &types.UserResponse{}

	err := s.db.QueryRow(
		query,
		user.Contact,
		user.Email,
		time.Now().UTC(),
		id,
	).Scan(
		&response.ID,
		&response.Contact,
		&response.Email,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
