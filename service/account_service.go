package service

import (
	"time"

	"github.com/Feggah/banking-api/domain"
	"github.com/Feggah/banking-api/dto"
	"github.com/Feggah/banking-api/errors"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, error)
	CreateTransaction(dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}

func (s DefaultAccountService) NewAccount(r dto.NewAccountRequest) (*dto.NewAccountResponse, error) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		CustomerID:  r.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: r.AccountType,
		Amount:      r.Amount,
		Status:      "1",
	}
	account, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := account.ToNewAccountResponseDto()
	return &response, nil
}

func (s DefaultAccountService) CreateTransaction(r dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {
	if err := r.ValidateRequest(); err != nil {
		return nil, err
	}

	if r.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(r.AccountID)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(r.Amount) {
			return nil, errors.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountID: r.AccountID,
		Amount:    r.Amount,
		Type:      r.Type,
		Date:      time.Now().Format("2006-01-02 15:04:05"),
	}
	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}

	response := transaction.ToDto()
	return response, nil
}
