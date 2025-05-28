package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/piyusharmap/go-banking/internal/middleware"
	"github.com/piyusharmap/go-banking/internal/types"
	"golang.org/x/crypto/bcrypt"
)

func (s *APIServer) HandleRegister(w http.ResponseWriter, r *http.Request) error {
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

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "registration successful",
		"user":    storedUser,
		"token":   token,
	})
}

func (s *APIServer) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "POST" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	loginRequest := types.AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		return err
	}

	defer r.Body.Close()

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

	storedUser.Password = ""

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "logged in",
		"email":   storedUser,
		"token":   token,
	})
}

func (s *APIServer) HandleUserUpdate(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "PUT" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}

	updateRequest := &types.UpdateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(updateRequest); err != nil {
		return err
	}

	defer r.Body.Close()

	user := &types.User{
		Contact: updateRequest.Contact,
		Email:   updateRequest.Email,
	}

	storedUser, err := s.Store.UpdateUser(id, user)

	if err != nil {
		return err
	}

	token, err := middleware.CreateJWT(user)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "update success",
		"user":    storedUser,
		"token":   token,
	})
}
