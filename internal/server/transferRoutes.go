package server

import (
	"encoding/json"
	"net/http"

	"github.com/piyusharmap/go-banking/internal/types"
	"github.com/piyusharmap/go-banking/internal/utility"
)

func (s *APIServer) HandleCreateTransfer(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "POST" {
		return ErrInvalidMethod()
	}

	request := &types.AmountTransferRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ErrInvalidRequest()
	}
	defer r.Body.Close()

	if request.SenderAccID == request.ReceiverAccID {
		return ErrInvalidRequest()
	}

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

func (s *APIServer) HandleTransferHistoryByAccount(w http.ResponseWriter, r *http.Request) error {
	requestMethod := r.Method

	if requestMethod != "GET" {
		return ErrInvalidMethod()
	}

	accountID, err := utility.GetRequestID(r)

	if err != nil {
		return ErrUnauthenticatedAccess()
	}

	customerID := r.Context().Value("customer_id").(int)

	if _, err := s.Store.GetAccountByID(accountID, customerID); err != nil {
		return ErrUnauthorizedAccess()
	}

	transfersResponse, err := s.Store.GetAllTransfer(accountID)

	if err != nil {
		return ErrInternalServer()
	}

	return WriteJSON(w, http.StatusOK, map[string]any{
		"message":   "user transfers",
		"transfers": transfersResponse,
	})
}
