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

	request := &types.Customer{}

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

	customer := &types.Customer{
		Contact:  request.Contact,
		Email:    request.Email,
		Password: string(password_hash),
	}

	response, err := s.Store.RegisterCustomer(customer)

	if err != nil {
		return ErrInternalServer()
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		// consider transactional rollback
		_, err := s.Store.DeleteCustomer(response.ID)

		if err != nil {
			return ErrInternalServer()
		}

		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusCreated, map[string]any{
		"message":  "customer registered",
		"customer": response,
		"token":    token,
	})
}

func (s *APIServer) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "POST" {
		return ErrInvalidMethod()
	}

	request := &types.Customer{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInvalidRequest()
	}
	defer r.Body.Close()

	if request.Contact == "" || request.Email == "" || request.Password == "" {
		return ErrInvalidRequest()
	}

	customer := &types.Customer{
		Contact:  request.Contact,
		Email:    request.Email,
		Password: request.Password,
	}

	storedCustomer, err := s.Store.GetCustomer(customer)

	if err != nil {
		return ErrUnauthenticatedAccess()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedCustomer.Password), []byte(customer.Password)); err != nil {
		return ErrInternalServer()
	}

	response := &types.CustomerResponse{
		ID:      storedCustomer.ID,
		Contact: storedCustomer.Contact,
		Email:   storedCustomer.Email,
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message":  "customer logged in",
		"customer": response,
		"token":    token,
	})
}

func (s *APIServer) HandleCustomerUpdate(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "PUT" {
		return ErrInvalidMethod()
	}

	id, err := utility.GetRequestID(r)

	if err != nil {
		return ErrInvalidRequest()
	}

	request := &types.UpdateCustomerRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInternalServer()
	}
	defer r.Body.Close()

	if request.Contact == "" || request.Email == "" {
		return ErrInvalidRequest()
	}

	CustomerID := r.Context().Value("customer_id").(int)

	if CustomerID != id {
		return ErrUnauthorizedAccess()
	}

	customer := &types.UpdateCustomerRequest{
		Contact: request.Contact,
		Email:   request.Email,
	}

	response, err := s.Store.UpdateCustomer(id, customer)

	if err != nil {
		return ErrInternalServer()
	}

	token, err := middleware.CreateJWT(response)

	if err != nil {
		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message":  "customer updated",
		"customer": response,
		"token":    token,
	})
}
