package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/piyusharmap/go-banking/internal/types"
	"github.com/piyusharmap/go-banking/internal/utility"
)

func (s *APIServer) HandleAccount(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	switch requestMethod {
	case "POST":
		return s.HandleCreateAccount(w, r)
	}

	return fmt.Errorf("invalid request method:%v", requestMethod)
}

func (s *APIServer) HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	request := &types.Account{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return err
	}

	defer r.Body.Close()

	authID := r.Context().Value("user_id").(int)

	if authID != request.UserID {
		return fmt.Errorf("permission denied")
	}

	account := &types.Account{
		UserID:    authID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Currency:  request.Currency,
	}

	response, err := s.Store.RegisterAccount(account)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "account registered",
		"account": response,
	})
}

func (s *APIServer) HandleAccountByID(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	switch requestMethod {
	case "GET":
		return s.HandleGetAccount(w, r)
	case "PUT":
		return s.HandleUpdateAccount(w, r)
	}

	return fmt.Errorf("invalid request method:%v", requestMethod)
}

func (s *APIServer) HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := utility.GetRequestID(r)

	if err != nil {
		return err
	}

	authID := r.Context().Value("user_id").(int)

	response, err := s.Store.GetAccountByID(id, authID)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"account": response,
	})
}

func (s *APIServer) HandleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := utility.GetRequestID(r)

	if err != nil {
		return err
	}

	request := &types.UpdateAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return err
	}

	if request.FirstName == "" || request.Currency == "" {
		return fmt.Errorf("first_name and currency can't be empty")
	}

	authID := r.Context().Value("user_id").(int)

	response, err := s.Store.UpdateAccount(id, authID, request)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "account updated",
		"account": response,
	})
}
