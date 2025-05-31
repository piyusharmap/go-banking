package storage

import (
	"time"

	"github.com/piyusharmap/go-banking/internal/types"
	"github.com/piyusharmap/go-banking/internal/utility"
)

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

func (s *PostgresStore) GetAccountByID(id int, authID int) (*types.AccountResponse, error) {
	query := `SELECT id, user_id, first_name, last_name, account_number
	FROM account
	WHERE id=$1 AND user_id=$2`

	response := &types.AccountResponse{}

	err := s.db.QueryRow(
		query,
		id,
		authID,
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

func (s *PostgresStore) UpdateAccount(id int, authID int, account *types.UpdateAccountRequest) (*types.AccountResponse, error) {
	query := `UPDATE account
	SET first_name=$1, last_name=$2, currency=$3, updated_at=$4 
	WHERE id=$5 AND user_id=$6
	RETURNING id, user_id, first_name, last_name, account_number`

	response := &types.AccountResponse{}

	err := s.db.QueryRow(
		query,
		account.FirstName,
		account.LastName,
		account.Currency,
		time.Now().UTC(),
		id,
		authID,
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
