package server

import (
	"encoding/json"
	"net/http"

	"github.com/piyusharmap/go-banking/internal/storage"
)

type APIServer struct {
	ListenAddr string
	Store      storage.Storage
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

type APIErrorResponse struct {
	Status  int
	Message string
}

func NewAPIServer(listenAddr string, store *storage.PostgresStore) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
		Store:      store,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (e *APIErrorResponse) Error() string {
	return e.Message
}

// error functions
func ErrUnauthenticatedAccess() *APIErrorResponse {
	return &APIErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: "access denied, unauthenticated access",
	}
}

func ErrUnauthorizedAccess() *APIErrorResponse {
	return &APIErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: "access denied, unauthorized access",
	}
}

func ErrInvalidRequest() *APIErrorResponse {
	return &APIErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "invalid request",
	}
}

func ErrInternalServer() *APIErrorResponse {
	return &APIErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
	}
}

func ErrInvalidMethod() *APIErrorResponse {
	return &APIErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "invalid request method",
	}
}
