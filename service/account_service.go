package service

import (
	"time"

	"github.com/Feggah/banking-api/domain"
	"github.com/Feggah/banking-api/dto"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, error)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
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

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
