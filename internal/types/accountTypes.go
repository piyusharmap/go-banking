package types

import "time"

type Account struct {
	CustomerID int    `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name,omitempty"`
}

type UpdateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type FetchBalanceRequest struct {
	ID            int    `json:"id"`
	AccountNumber string `json:"account_number"`
}

type AccountModel struct {
	ID            int
	CustomerID    int
	AccountNumber string
	Balance       int64
	CreatedAt     time.Time
}

type AccountResponse struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber string `json:"account_number"`
}

type BalanceQueryResponse struct {
	ID            int    `json:"id"`
	AccountNumber string `json:"account_number"`
	Balance       string `json:"balance"`
}
