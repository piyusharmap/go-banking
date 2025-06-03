package server

import (
	"encoding/json"
	"net/http"

	"github.com/piyusharmap/go-banking/internal/middleware"
	"github.com/piyusharmap/go-banking/internal/types"
	"github.com/piyusharmap/go-banking/internal/utility"
	"golang.org/x/crypto/bcrypt"
)

func (s *APIServer) HandleRegister(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "POST" {
		return ErrInvalidMethod()
	}

	request := &types.User{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInvalidRequest()
	}
	defer r.Body.Close()

	if request.Contact == "" || request.Email == "" || request.Password == "" {
		return ErrInvalidRequest()
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return ErrUnauthenticatedAccess()
	}

	user := &types.User{
		Contact:  request.Contact,
		Email:    request.Email,
		Password: string(password_hash),
	}

	response, err := s.Store.RegisterUser(user)

	if err != nil {
		return ErrInternalServer()
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		// consider transactional rollback
		_, err := s.Store.DeleteUser(response.ID)

		if err != nil {
			return ErrInternalServer()
		}

		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusCreated, map[string]any{
		"message": "user registered",
		"user":    response,
		"token":   token,
	})
}

func (s *APIServer) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "POST" {
		return ErrInvalidMethod()
	}

	request := &types.User{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInvalidRequest()
	}
	defer r.Body.Close()

	if request.Contact == "" || request.Email == "" || request.Password == "" {
		return ErrInvalidRequest()
	}

	user := &types.User{
		Contact:  request.Contact,
		Email:    request.Email,
		Password: request.Password,
	}

	storedUser, err := s.Store.GetUser(user)

	if err != nil {
		return ErrUnauthenticatedAccess()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return ErrInternalServer()
	}

	response := &types.UserResponse{
		ID:      storedUser.ID,
		Contact: storedUser.Contact,
		Email:   storedUser.Email,
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		return ErrInternalServer()
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
		return ErrInvalidMethod()
	}

	id, err := utility.GetRequestID(r)

	if err != nil {
		return ErrInvalidRequest()
	}

	request := &types.UpdateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInternalServer()
	}

	defer r.Body.Close()

	userID := r.Context().Value("user_id").(int)

	if userID != id {
		return ErrUnauthorizedAccess()
	}

	user := &types.UpdateUserRequest{
		Contact: request.Contact,
		Email:   request.Email,
	}

	response, err := s.Store.UpdateUser(id, user)

	if err != nil {
		return ErrInternalServer()
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "user updated",
		"user":    response,
		"token":   token,
	})
}
