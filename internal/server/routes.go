package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyusharmap/go-banking/internal/types"

	"github.com/piyusharmap/go-banking/internal/middleware"

	"golang.org/x/crypto/bcrypt"
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
	requestMethod := r.Method

	if requestMethod != "POST" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	registerRequest := types.AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(&registerRequest)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	encrypted_password, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	fmt.Println(string(encrypted_password))

	token, err := middleware.CreateJWT(&registerRequest)

	if err != nil {
		return err
	}

	fmt.Println(token)

	WriteJSON(w, http.StatusOK, registerRequest)

	return nil
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {

	// take user input for email/contact, password and jwt
	// will validate the jwt token and password
	// send back user details and jwt in reponse

	requestMethod := r.Method

	if requestMethod != "POST" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	loginRequest := types.AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		return err
	}

	// if err := bcrypt.CompareHashAndPassword([]byte(encrypted_pass), []byte(loginRequest.Password)); err != nil {
	// 	return err
	// }

	WriteJSON(w, http.StatusOK, loginRequest)

	return nil
}
