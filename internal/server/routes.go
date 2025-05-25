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

	log.Println("server is running on port:", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)
}

// auth routes

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {

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

	user := &types.User{
		Contact:  registerRequest.Contact,
		Email:    registerRequest.Email,
		Password: string(encrypted_password),
	}

	token, err := middleware.CreateJWT(user)

	if err != nil {
		return err
	}

	storedUser, err := s.Store.RegisterUser(user)

	if err != nil {
		return err
	}

	WriteJSON(w, http.StatusOK, map[string]any{
		"contact": storedUser.Contact,
		"email":   storedUser.Email,
		"token":   token,
	})

	return nil
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {

	requestMethod := r.Method

	if requestMethod != "POST" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	loginRequest := types.AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		return err
	}

	user := &types.User{
		Contact:  loginRequest.Contact,
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	storedUser, err := s.Store.GetUser(user)

	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginRequest.Password)); err != nil {
		return err
	}

	token, err := middleware.CreateJWT(user)

	if err != nil {
		return err
	}

	WriteJSON(w, http.StatusOK, map[string]any{
		"contact": storedUser.Contact,
		"email":   storedUser.Email,
		"token":   token,
	})

	return nil
}
