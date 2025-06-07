package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)

		if err != nil {
			if apiError, ok := err.(*APIErrorResponse); ok {
				WriteJSON(w, apiError.Status, &APIError{
					Error: apiError.Message,
				})
			} else {
				WriteJSON(w, http.StatusInternalServerError, &APIError{
					Error: err.Error(),
				})
			}
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// customer routes
	router.HandleFunc("/customer/register", makeHTTPHandleFunc(s.HandleRegister))
	router.HandleFunc("/customer/login", makeHTTPHandleFunc(s.HandleLogin))
	router.HandleFunc("/customer/{id}", withJWTAuth(makeHTTPHandleFunc(s.HandleCustomerUpdate), s))

	// account routes
	router.HandleFunc("/account", withJWTAuth(makeHTTPHandleFunc(s.HandleAccount), s))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.HandleAccountByID), s))

	// transfer routes
	router.HandleFunc("/transfer", withJWTAuth(makeHTTPHandleFunc(s.HandleTransfer), s))

	log.Println("server is running on port:", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)
}
