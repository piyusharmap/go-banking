package server

import (
	"fmt"
	"net/http"
)

func (s *APIServer) HandleAccount(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	switch requestMethod {
	case "POST":
		return s.HandleCreateAccount(w, r)
	case "GET":
		return s.HandleGetAccount(w, r)
	}

	return fmt.Errorf("invalid request method:%v", requestMethod)
}

func (s *APIServer) HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) HandleAccountByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
