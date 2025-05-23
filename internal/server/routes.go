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

	router.HandleFunc("/auth/register", makeHTTPHandleFunc(s.handleRegister))
	router.HandleFunc("/auth/login", makeHTTPHandleFunc(s.handleLogin))

	log.Println("server is running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// auth routes

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {

	// take user input for contact, email and password
	// encrypt the password and store in the database
	// create a jwt token based on email and contact and store in database
	// send back user details and jwt token in response

	WriteJSON(w, http.StatusOK, "register route")

	return nil
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {

	// take user input for email/contact, password and jwt
	// will validate the jwt token and password
	// send back user details and jwt in reponse

	WriteJSON(w, http.StatusOK, "login route")

	return nil
}
