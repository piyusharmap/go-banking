package server

import (
	"context"
	"net/http"

	"github.com/piyusharmap/go-banking/internal/middleware"
)

// check if the request header contains jwt token, throw error if fails
// validate the tokeni, throw error if fails
// extract the claims, throw error if fails
// match claims with corresponding DB entry, throw error if fails
// pass down the request with context (logged in user ID)
func withJWTAuth(handlerFunc http.HandlerFunc, s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			err := ErrUnauthenticatedAccess()

			WriteJSON(w, err.Status, &APIError{
				Error: err.Message,
			})

			return
		}

		token, err := middleware.ValidateJWT(tokenString)

		if err != nil {
			err := ErrUnauthenticatedAccess()

			WriteJSON(w, err.Status, &APIError{
				Error: err.Message,
			})

			return
		}

		claims, ok := token.Claims.(*middleware.CustomJWTClaims)

		if !ok {
			err := ErrUnauthenticatedAccess()

			WriteJSON(w, err.Status, &APIError{
				Error: err.Message,
			})

			return
		}

		user, err := s.Store.GetUserByID(claims.ID)

		if err != nil {
			err := ErrUnauthenticatedAccess()

			WriteJSON(w, err.Status, &APIError{
				Error: err.Message,
			})

			return
		}

		if claims.Contact != user.Contact || claims.Email != user.Email {
			err := ErrUnauthenticatedAccess()

			WriteJSON(w, err.Status, &APIError{
				Error: err.Message,
			})

			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.ID)

		handlerFunc(w, r.WithContext(ctx))
	}
}
