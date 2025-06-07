package server

import "net/http"

func (s *APIServer) HandleTransfer(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	switch requestMethod {
	case "POST":
		return s.HandleCreateAccount(w, r)
	}

	return ErrInvalidRequest()
}

func (s *APIServer) HandleCreateTransfer(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, "Transer Route")
}
