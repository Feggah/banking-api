package service

import (
	"github.com/Feggah/banking-api/domain"
	"github.com/Feggah/banking-api/dto"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, error)
	GetCustomer(string) (*dto.CustomerResponse, error)
	GetAllCustomersByStatus(string) ([]dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, error) {
	c, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	dtos := make([]dto.CustomerResponse, 0)
	for _, customer := range c {
		dtos = append(dtos, *customer.ToDto())
	}

	return dtos, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, error) {
	c, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return c.ToDto(), nil
}

func (s DefaultCustomerService) GetAllCustomersByStatus(status string) ([]dto.CustomerResponse, error) {
	c, err := s.repo.GetAllCustomersByStatus(status)
	if err != nil {
		return nil, err
	}

	dtos := make([]dto.CustomerResponse, 0)
	for _, customer := range c {
		dtos = append(dtos, *customer.ToDto())
	}

	return dtos, nil
}

func NewCustomerService(r domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: r}
}
