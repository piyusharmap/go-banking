package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/piyusharmap/go-banking/internal/middleware"
	"github.com/piyusharmap/go-banking/internal/types"
	"github.com/piyusharmap/go-banking/internal/utility"
	"golang.org/x/crypto/bcrypt"
)

func (s *APIServer) HandleRegister(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "POST" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	request := types.User{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return err
	}

	// closing body to prevent resourse leak
	defer r.Body.Close()

	password_hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := &types.User{
		Contact:  request.Contact,
		Email:    request.Email,
		Password: string(password_hash),
	}

	response, err := s.Store.RegisterUser(user)

	if err != nil {
		return err
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "user registered",
		"user":    response,
		"token":   token,
	})
}

func (s *APIServer) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "POST" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	request := types.User{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	user := &types.User{
		Contact:  request.Contact,
		Email:    request.Email,
		Password: request.Password,
	}

	storedUser, err := s.Store.GetUser(user)

	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return err
	}

	response := &types.UserResponse{
		ID:      storedUser.ID,
		Contact: storedUser.Contact,
		Email:   storedUser.Email,
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "user logged in",
		"user":    response,
		"token":   token,
	})
}

func (s *APIServer) HandleUserUpdate(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "PUT" {
		return fmt.Errorf("invalid request method:%v", requestMethod)
	}

	id, err := utility.GetRequestID(r)

	if err != nil {
		return err
	}

	request := &types.UpdateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return err
	}

	defer r.Body.Close()

	user := &types.UpdateUserRequest{
		Contact: request.Contact,
		Email:   request.Email,
	}

	response, err := s.Store.UpdateUser(id, user)

	if err != nil {
		return err
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "user updated",
		"user":    response,
		"token":   token,
	})
}
