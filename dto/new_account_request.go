package dto

import (
	"strings"

	"github.com/Feggah/banking-api/errors"
)

type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() error {
	if r.Amount < 5000 {
		return errors.NewValidationError("To open a new account you need to deposit at least 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errors.NewValidationError("Account type should be 'checking' or 'saving'")
	}
	return nil
}
