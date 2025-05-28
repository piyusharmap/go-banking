package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/piyusharmap/go-banking/internal/middleware"
)

func withJWTAuth(handlerFunc http.HandlerFunc, s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		token, err := middleware.VerifyJWT(tokenString)

		if err != nil {
			ThrowPermissionDenied(w)
			return
		}

		idStr := mux.Vars(r)["id"]

		id, err := strconv.Atoi(idStr)

		if err != nil {
			ThrowPermissionDenied(w)
			return
		}

		user, err := s.Store.GetUserByID(id)

		if err != nil {
			ThrowPermissionDenied(w)
			return
		}

		if claims, ok := token.Claims.(*middleware.CustomJWTClaims); ok {
			if claims.Contact != user.Contact && claims.Email != user.Email {
				ThrowPermissionDenied(w)
				return
			}
		} else {
			ThrowPermissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}
