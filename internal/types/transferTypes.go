package types

import "time"

type TransferStage string

const (
	TransferStagePending TransferStage = "PENDING"
	TransferStageSuccess TransferStage = "SUCCESS"
	TransferStageFailed  TransferStage = "FAILED"
)

type AmountTransfer struct {
	SenderAccID   int           `json:"sender_account_id"`
	ReceiverAccID int           `json:"receiver_account_id"`
	Amount        int64         `json:"amount"`
	Stage         TransferStage `json:"transfer_stage"`
	Remark        string        `json:"remark,omitempty"`
}

type AmountTransferRequest struct {
	SenderAccID   int    `json:"sender_account_id"`
	ReceiverAccID int    `json:"receiver_account_id"`
	Amount        int64  `json:"amount"`
	Remark        string `json:"remark,omitempty"`
}

type AmountTransferModel struct {
	ID            int
	SenderAccID   int
	ReceiverAccID int
	Amount        int64
	Stage         TransferStage
	Remark        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type AmountTransferResponse struct {
	ID                int           `json:"id"`
	SenderAccID       int           `json:"sender_account_id"`
	SenderAccNumber   string        `json:"sender_account_number"`
	ReceiverAccID     int           `json:"receiver_account_id"`
	ReceiverAccNumber string        `json:"receiver_account_number"`
	Amount            int64         `json:"amount"`
	Stage             TransferStage `json:"transfer_stage"`
	Remark            string        `json:"remark,omitempty"`
	CreatedAt         time.Time     `json:"created_at"`
}
