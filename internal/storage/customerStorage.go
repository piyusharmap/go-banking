package storage

import (
	"time"

	"github.com/piyusharmap/go-banking/internal/types"
)

func (s *PostgresStore) RegisterCustomer(customer *types.Customer) (*types.CustomerResponse, error) {
	query := `INSERT INTO customer (contact, email, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, contact, email`

	response := &types.CustomerResponse{}

	err := s.db.QueryRow(
		query,
		customer.Contact,
		customer.Email,
		customer.Password,
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

func (s *PostgresStore) GetCustomer(customer *types.Customer) (*types.CustomerModel, error) {
	query := `SELECT id, contact, email, password_hash
	From customer
	WHERE contact=$1 AND email=$2`

	storedCustomer := &types.CustomerModel{}

	err := s.db.QueryRow(
		query,
		customer.Contact,
		customer.Email,
	).Scan(
		&storedCustomer.ID,
		&storedCustomer.Contact,
		&storedCustomer.Email,
		&storedCustomer.Password,
	)

	if err != nil {
		return nil, err
	}

	return storedCustomer, nil
}

func (s *PostgresStore) GetCustomerByID(id int) (*types.CustomerResponse, error) {
	query := `SELECT id, contact, email
	From customer
	WHERE id=$1`

	response := &types.CustomerResponse{}

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

func (s *PostgresStore) UpdateCustomer(id int, customer *types.UpdateCustomerRequest) (*types.CustomerResponse, error) {
	query := `UPDATE customer
	SET contact=$1, email=$2, updated_at=$3
	WHERE id=$4
	RETURNING id, contact, email`

	response := &types.CustomerResponse{}

	err := s.db.QueryRow(
		query,
		customer.Contact,
		customer.Email,
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

func (s *PostgresStore) DeleteCustomer(id int) (*types.CustomerResponse, error) {
	query := `DELETE From customer
	WHERE id=$1
	RETURNING id, contact, email`

	response := &types.CustomerResponse{}

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
