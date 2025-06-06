package types

import "time"

type Customer struct {
	Contact  string `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateCustomerRequest struct {
	Contact string `json:"contact"`
	Email   string `json:"email"`
}

type CustomerModel struct {
	ID        int
	Contact   string
	Email     string
	Password  string
	CreatedAt time.Time
	updatedAt time.Time
}

type CustomerResponse struct {
	ID      int    `json:"id"`
	Contact string `json:"contact"`
	Email   string `json:"email"`
}
