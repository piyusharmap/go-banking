package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)

		if err != nil {
			WriteJSON(w, http.StatusBadRequest, &APIError{
				Error: err.Error(),
			})
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// auth routes
	router.HandleFunc("/auth/register", makeHTTPHandleFunc(s.HandleRegister))
	router.HandleFunc("/auth/login", makeHTTPHandleFunc(s.HandleLogin))
	router.HandleFunc("/user/{id}", withJWTAuth(makeHTTPHandleFunc(s.HandleUserUpdate), s))

	// account routes
	router.HandleFunc("/account", makeHTTPHandleFunc(s.HandleAccount))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.HandleAccountByID), s))

	log.Println("server is running on port:", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)
}

func ThrowPermissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusBadRequest, &APIError{
		Error: "permission denied",
	})
}
