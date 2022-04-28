package domain

import "github.com/Feggah/banking-api/dto"

type Transaction struct {
	ID        string  `db:"transaction_id"`
	AccountID string  `db:"account_id"`
	Amount    float64 `db:"amount"`
	Type      string  `db:"transaction_type"`
	Date      string  `db:"trasaction_date"`
}

func (t *Transaction) ToDto() *dto.CreateTransactionResponse {
	return &dto.CreateTransactionResponse{
		TransactionID: t.ID,
		Balance:       t.Amount,
	}
}

func (t *Transaction) IsWithdrawal() bool {
	return t.Type == "withdrawal"
}

type TransactionRepository interface {
	Create(Transaction) (*Transaction, error)
}
