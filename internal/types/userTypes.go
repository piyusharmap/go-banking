package types

import "time"

type User struct {
	Contact  string `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Contact string `json:"contact"`
	Email   string `json:"email"`
}

type UserModel struct {
	ID        int
	Contact   string
	Email     string
	Password  string
	CreatedAt time.Time
	updatedAt time.Time
}

type UserResponse struct {
	ID      int    `json:"id"`
	Contact string `json:"contact"`
	Email   string `json:"email"`
}
