package server

import (
	"encoding/json"
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

	return ErrInvalidMethod()
}

func (s *APIServer) HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	request := &types.Account{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInvalidRequest()
	}
	defer r.Body.Close()

	if request.FirstName == "" {
		return ErrInvalidRequest()
	}

	CustomerID := r.Context().Value("customer_id").(int)

	if CustomerID != request.CustomerID {
		return ErrUnauthorizedAccess()
	}

	account := &types.Account{
		CustomerID: CustomerID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
	}

	response, err := s.Store.RegisterAccount(account)

	if err != nil {
		return ErrInternalServer()
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
	case "DELETE":
		return s.HandleRemoveAccount(w, r)
	}

	return ErrInvalidMethod()
}

func (s *APIServer) HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := utility.GetRequestID(r)

	if err != nil {
		return ErrUnauthenticatedAccess()
	}

	CustomerID := r.Context().Value("customer_id").(int)

	response, err := s.Store.GetAccountByID(id, CustomerID)

	if err != nil {
		return ErrUnauthorizedAccess()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"account": response,
	})
}

func (s *APIServer) HandleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := utility.GetRequestID(r)

	if err != nil {
		return ErrUnauthenticatedAccess()
	}

	request := &types.UpdateAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInvalidRequest()
	}

	if request.FirstName == "" {
		return ErrInvalidRequest()
	}

	CustomerID := r.Context().Value("customer_id").(int)

	response, err := s.Store.UpdateAccount(id, CustomerID, request)

	if err != nil {
		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message": "account updated",
		"account": response,
	})
}

func (s *APIServer) HandleRemoveAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := utility.GetRequestID(r)

	if err != nil {
		return ErrInvalidRequest()
	}

	CustomerID := r.Context().Value("customer_id").(int)

	response, err := s.Store.RemoveAccount(id, CustomerID)

	if err != nil {
		return ErrUnauthorizedAccess()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"account": response,
		"message": "account removed",
	})
}
