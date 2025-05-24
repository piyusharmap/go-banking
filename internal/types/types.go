package types

type User struct {
	Contact  int64  `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Contact  int64  `json:"contact"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateNewUser(contact int64, email, password string) *User {
	return nil
}
