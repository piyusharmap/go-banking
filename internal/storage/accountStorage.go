package storage

import (
	"time"

	"github.com/piyusharmap/go-banking/internal/types"
	"github.com/piyusharmap/go-banking/internal/utility"
)

func (s *PostgresStore) GetCustomerAccounts(customerID int) ([]*types.AccountResponse, error) {
	query := `SELECT id, customer_id, first_name, last_name, account_nummber
	FROM account
	WHERE customer_id=$1`

	rows, err := s.db.Query(query, customerID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responseAccounts []*types.AccountResponse

	for rows.Next() {
		var account types.AccountResponse

		err := rows.Scan(
			&account.ID,
			&account.CustomerID,
			&account.FirstName,
			&account.LastName,
			&account.AccountNumber,
		)

		if err != nil {
			return nil, err
		}

		responseAccounts = append(responseAccounts, &account)
	}

	return responseAccounts, nil
}

func (s *PostgresStore) RegisterAccount(account *types.Account) (*types.AccountResponse, error) {
	query := `INSERT INTO account (customer_id, first_name, last_name, account_number)
	VALUES ($1, $2, $3, $4)
	RETURNING id, customer_id, first_name, last_name, account_number`

	response := &types.AccountResponse{}

	accNum := utility.GenerateAccNumber()

	err := s.db.QueryRow(
		query,
		account.CustomerID,
		account.FirstName,
		account.LastName,
		accNum,
	).Scan(
		&response.ID,
		&response.CustomerID,
		&response.FirstName,
		&response.LastName,
		&response.AccountNumber,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostgresStore) GetAccountByID(id, customerID int) (*types.AccountResponse, error) {
	query := `SELECT id, customer_id, first_name, last_name, account_number
	FROM account
	WHERE id=$1 AND customer_id=$2`

	response := &types.AccountResponse{}

	err := s.db.QueryRow(
		query,
		id,
		customerID,
	).Scan(
		&response.ID,
		&response.CustomerID,
		&response.FirstName,
		&response.LastName,
		&response.AccountNumber,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostgresStore) UpdateAccount(id, CustomerID int, account *types.UpdateAccountRequest) (*types.AccountResponse, error) {
	query := `UPDATE account
	SET first_name=$1, last_name=$2, updated_at=$3
	WHERE id=$4 AND customer_id=$5
	RETURNING id, customer_id, first_name, last_name, account_number`

	response := &types.AccountResponse{}

	err := s.db.QueryRow(
		query,
		account.FirstName,
		account.LastName,
		time.Now().UTC(),
		id,
		CustomerID,
	).Scan(
		&response.ID,
		&response.CustomerID,
		&response.FirstName,
		&response.LastName,
		&response.AccountNumber,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostgresStore) RemoveAccount(id, CustomerID int) (*types.AccountResponse, error) {
	query := `DELETE FROM account
	WHERE id=$1 AND customer_id=$2
	RETURNING id, customer_id, first_name, last_name, account_number`

	response := &types.AccountResponse{}

	err := s.db.QueryRow(
		query,
		id,
		CustomerID,
	).Scan(
		&response.ID,
		&response.CustomerID,
		&response.FirstName,
		&response.LastName,
		&response.AccountNumber,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
