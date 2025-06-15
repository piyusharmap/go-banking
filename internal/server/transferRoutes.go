package server

import (
	"encoding/json"
	"net/http"

	"github.com/piyusharmap/go-banking/internal/types"
)

func (s *APIServer) HandleTransfer(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	switch requestMethod {
	case "POST":
		return s.HandleCreateAccount(w, r)
	}

	return ErrInvalidRequest()
}

func (s *APIServer) HandleCreateTransfer(w http.ResponseWriter, r *http.Request) error {
	request := &types.AmountTransferRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInvalidRequest()
	}
	defer r.Body.Close()

	customerID := r.Context().Value("customer_id").(int)

	customerBalance, err := s.Store.FetchRawBalance(request.SenderAccID, customerID)

	if err != nil {
		return ErrUnauthorizedAccess()
	}

	if customerBalance < request.Amount {
		return WriteJSON(w, http.StatusConflict, map[string]any{
			"message": "can't proceed with transfer, funds not enough",
		})
	}

	amountTransfer := &types.AmountTransfer{
		SenderAccID:   request.SenderAccID,
		ReceiverAccID: request.ReceiverAccID,
		Amount:        request.Amount,
		Stage:         "PENDING",
		Remark:        request.Remark,
	}

	response, err := s.Store.RegisterTransfer(amountTransfer)

	if err != nil {
		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message":  "transfer status",
		"transfer": response,
	})
}
