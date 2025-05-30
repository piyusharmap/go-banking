package server

import (
	"net/http"

	"github.com/piyusharmap/go-banking/internal/middleware"
)

func withJWTAuth(handlerFunc http.HandlerFunc, s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			ThrowPermissionDenied(w)
			return
		}

		token, err := middleware.ValidateJWT(tokenString)

		if err != nil {
			ThrowPermissionDenied(w)
			return
		}

		claims, ok := token.Claims.(*middleware.CustomJWTClaims)

		if !ok {
			ThrowPermissionDenied(w)
			return
		}

		user, err := s.Store.GetUserByID(claims.ID)

		if err != nil {
			ThrowPermissionDenied(w)
			return
		}

		if claims.Contact != user.Contact || claims.Email != user.Email {
			ThrowPermissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}
