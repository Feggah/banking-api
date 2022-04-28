package dto

type CreateTransactionResponse struct {
	Balance       float64 `json:"balance"`
	TransactionID string  `json:"transaction_id"`
}
