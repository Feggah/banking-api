package domain

import "github.com/Feggah/banking-api/dto"

type Customer struct {
	ID        string `db:"customerId"`
	Name      string `db:"name"`
	City      string `db:"city"`
	ZIPCode   string `db:"zipCode"`
	BirthDate string `db:"birthDate"`
	Status    string `db:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c *Customer) ToDto() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		ID:        c.ID,
		Name:      c.Name,
		City:      c.City,
		ZIPCode:   c.ZIPCode,
		BirthDate: c.BirthDate,
		Status:    c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetById(string) (*Customer, error)
	GetAllCustomersByStatus(string) ([]Customer, error)
}
