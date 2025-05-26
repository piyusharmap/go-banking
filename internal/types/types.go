package types

import "time"

type User struct {
	Contact  int64  `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StoredUser struct {
	Id       int    `json:"id"`
	Contact  int64  `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Account struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber string `json:"account_number"`
	Balance       int64  `json:"balance"`
	UserId        int    `json:"user_id"`
}

type StoredAccount struct {
	Id            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	AccountNumber string    `json:"account_number"`
	Balance       int64     `json:"balance"`
	UserId        int       `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AuthRequest struct {
	Contact  int64  `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Contact int64  `json:"contact"`
	Email   string `json:"email"`
}

func CreateNewUser(contact int64, email, password string) *User {
	return nil
}
