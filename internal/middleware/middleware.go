package middleware

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/piyusharmap/go-banking/internal/types"
)

type CustomJWTClaims struct {
	Contact int64  `json:"contact"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJWT(user *types.User) (string, error) {
	claims := CustomJWTClaims{
		user.Contact,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
}
