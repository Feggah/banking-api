package dto

import "github.com/Feggah/banking-api/errors"

type CreateTransactionRequest struct {
	Type       string  `json:"transaction_type"`
	Amount     float64 `json:"amount"`
	AccountID  string  `json:"account_id"`
	CustomerID string  `json:"customer_id"`
}

const (
	WithdrawalType = "withdrawal"
	DepositType    = "deposit"
)

func (c *CreateTransactionRequest) ValidateRequest() error {
	if c.Amount < 0 {
		return errors.NewValidationError("Amount cannot be negative")
	}

	if c.Type != WithdrawalType && c.Type != DepositType {
		return errors.NewValidationError("Type must be 'withdrawal' or 'deposit'")
	}
	return nil
}

func (c *CreateTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return c.Type == WithdrawalType
}
