package middleware

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/piyusharmap/go-banking/internal/types"
)

type CustomJWTClaims struct {
	ID      int    `json:"id"`
	Contact string `json:"contact"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

// to create jwt token based on id, contact and email
func CreateJWT(customer *types.CustomerResponse) (string, error) {
	claims := CustomJWTClaims{
		customer.ID,
		customer.Contact,
		customer.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// to validate the token provided by customer based on custom claims
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
}
