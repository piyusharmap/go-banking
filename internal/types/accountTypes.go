package types

import "time"

type Account struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Currency  string `json:"currency,omitempty"`
}

type UpdateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Currency  string `json:"currency"`
}

type AccountModel struct {
	ID            int
	UserID        int
	AccountNumber string
	Balance       int64
	Currency      string
	CreatedAt     time.Time
}

type AccountResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber string `json:"account_number"`
}
