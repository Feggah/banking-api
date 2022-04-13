package domain

import (
	"database/sql"

	"github.com/Feggah/banking-api/errors"
	"github.com/Feggah/banking-api/logger"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "SELECT customerId, name, city, zipCode, birthDate, status FROM customers"
	customers := make([]Customer, 0)
	err := c.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

func (c CustomerRepositoryDb) GetById(id string) (*Customer, error) {
	customerSQL := "SELECT customerId, name, city, zipCode, birthDate, status FROM customers WHERE customerId = ?"

	var customer Customer
	err := c.client.Get(&customer, customerSQL, id)
	if err != nil {
		logger.Error("Error while scanning customer " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		}
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return &customer, nil
}

func (c CustomerRepositoryDb) GetAllCustomersByStatus(status string) ([]Customer, error) {
	if status == "active" {
		status = "1"
	} else {
		status = "0"
	}
	customers := make([]Customer, 0)
	customerSQL := "SELECT customerId, name, city, zipCode, birthDate, status FROM customers WHERE status = ?"
	err := c.client.Select(&customers, customerSQL, status)
	if err != nil {
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

func NewCustomerRepositoryDb(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: client}
}
