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
