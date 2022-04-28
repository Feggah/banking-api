package domain

import "github.com/Feggah/banking-api/dto"

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a *Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}

func (a *Account) CanWithdraw(amount float64) bool {
	return amount < a.Amount
}

type AccountRepository interface {
	Save(Account) (*Account, error)
	FindBy(string) (*Account, error)
	SaveTransaction(Transaction) (*Transaction, error)
}
