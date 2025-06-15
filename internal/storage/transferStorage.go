package storage

import (
	"fmt"

	"github.com/piyusharmap/go-banking/internal/types"
)

// first intiate a transfer
// fetch sender account and lock the account
// fetch receiver account
// deduct amount from sender account and credit in receiver account
// create entry in transfer table
// commit the transaction
func (s *PostgresStore) RegisterTransfer(amountTransfer *types.AmountTransfer) (*types.AmountTransferResponse, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query1 := `SELECT account_number, balance
	FROM account
	WHERE id=$1
	FOR UPDATE`

	var senderAccNumber string
	var senderBalance int64

	if err := tx.QueryRow(
		query1,
		amountTransfer.SenderAccID,
	).Scan(
		&senderAccNumber,
		&senderBalance,
	); err != nil {
		return nil, err
	}

	if senderBalance < amountTransfer.Amount {
		return nil, fmt.Errorf("insufficient funds")
	}

	query2 := `SELECT account_number
	FROM account
	WHERE id=$1
	FOR UPDATE`

	var receiverAccNumber string

	if err := tx.QueryRow(
		query2,
		amountTransfer.ReceiverAccID,
	).Scan(
		&receiverAccNumber,
	); err != nil {
		return nil, err
	}

	query3 := `UPDATE account
	SET balance=balance-$1, updated_at=CURRENT_TIMESTAMP
	WHERE id=$2`

	if _, err := tx.Exec(
		query3,
		amountTransfer.Amount,
		amountTransfer.SenderAccID,
	); err != nil {
		return nil, err
	}

	query4 := `UPDATE account
	SET balance=balance + $1, updated_at=CURRENT_TIMESTAMP
	WHERE id=$2`

	if _, err := tx.Exec(
		query4,
		amountTransfer.Amount,
		amountTransfer.ReceiverAccID,
	); err != nil {
		return nil, err
	}

	query5 := `INSERT INTO amount_transfer(sender_account_id, receiver_account_id, amount, stage, remark)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, sender_account_id, receiver_account_id, amount, stage, remark, created_at`

	response := &types.AmountTransferResponse{}

	if err := tx.QueryRow(
		query5,
		amountTransfer.SenderAccID,
		amountTransfer.ReceiverAccID,
		amountTransfer.Amount,
		"COMPLETED",
		amountTransfer.Remark,
	).Scan(
		&response.ID,
		&response.SenderAccID,
		&response.ReceiverAccID,
		&response.Amount,
		&response.Stage,
		&response.Remark,
		&response.CreatedAt,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	response.SenderAccNumber = senderAccNumber
	response.ReceiverAccNumber = receiverAccNumber

	return response, nil
}

func (s *PostgresStore) GetAllTransfer(accountID int) ([]*types.AmountTransferResponse, error) {
	query := `SELECT 
		at.id,
		at.sender_account_id,
		sender.account_number AS sender_account_number,
		at.receiver_account_id,
		receiver.account_number AS receiver_account_number,
		at.amount,
		at.stage,
		at.remark,
		at.created_at
	FROM amount_transfer at
	JOIN account sender ON at.sender_account_id=sender.id
	JOIN account receiver ON at.receiver_account_id=receiver.id
	WHERE sender.id=$1 OR receiver.id=$1
	ORDER BY at.created_at DESC`

	rows, err := s.db.Query(query, accountID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responseTransfers []*types.AmountTransferResponse

	for rows.Next() {
		var transfer types.AmountTransferResponse

		err := rows.Scan(
			&transfer.ID,
			&transfer.SenderAccID,
			&transfer.SenderAccNumber,
			&transfer.ReceiverAccID,
			&transfer.ReceiverAccNumber,
			&transfer.Amount,
			&transfer.Stage,
			&transfer.Remark,
			&transfer.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		responseTransfers = append(responseTransfers, &transfer)
	}

	return responseTransfers, nil
}
