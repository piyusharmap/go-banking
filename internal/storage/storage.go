package storage

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/piyusharmap/go-banking/internal/types"
	"github.com/piyusharmap/go-banking/internal/utility"
)

type Storage interface {
	// user methods
	RegisterUser(*types.User) (*types.UserResponse, error)
	GetUser(*types.User) (*types.UserModel, error)
	GetUserByID(int) (*types.UserResponse, error)
	UpdateUser(int, *types.UpdateUserRequest) (*types.UserResponse, error)

	// account methods
	RegisterAccount(*types.Account) (*types.AccountResponse, error)
	GetAccountByID(int) (*types.AccountResponse, error)
	UpdateAccount(int, *types.UpdateAccountRequest) (*types.AccountResponse, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	if err := CreateUserTable(s.db); err != nil {
		return err
	}

	if err := CreateAccountTable(s.db); err != nil {
		return err
	}

	return nil
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

func (s *PostgresStore) UpdateUser(id int, user *types.UpdateUserRequest) (*types.UserResponse, error) {
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

func (s *PostgresStore) RegisterAccount(account *types.Account) (*types.AccountResponse, error) {
	query := `INSERT INTO account (user_id, first_name, last_name, account_number, currency)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, user_id, first_name, last_name, account_number`

	response := &types.AccountResponse{}

	accNum := utility.GenerateAccNumber()

	err := s.db.QueryRow(
		query,
		account.UserID,
		account.FirstName,
		account.LastName,
		accNum,
		account.Currency,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.FirstName,
		&response.LastName,
		&response.AccountNumber,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*types.AccountResponse, error) {
	query := `SELECT id, user_id, first_name, last_name, account_number
	FROM account
	WHERE id=$1`

	response := &types.AccountResponse{}

	err := s.db.QueryRow(
		query,
		id,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.FirstName,
		&response.LastName,
		&response.AccountNumber,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostgresStore) UpdateAccount(id int, account *types.UpdateAccountRequest) (*types.AccountResponse, error) {
	query := `UPDATE account
	SET first_name=$1, last_name=$2, currency=$3, updated_at=$4 
	WHERE id=$5
	RETURNING id, user_id, first_name, last_name, account_number`

	response := &types.AccountResponse{}

	err := s.db.QueryRow(
		query,
		account.FirstName,
		account.LastName,
		account.Currency,
		time.Now().UTC(),
		id,
	).Scan(
		&response.ID,
		&response.UserID,
		&response.FirstName,
		&response.LastName,
		&response.AccountNumber,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
