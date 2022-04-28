package domain

import (
	"database/sql"
	"strconv"

	"github.com/Feggah/banking-api/errors"
	"github.com/Feggah/banking-api/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}

func (d AccountRepositoryDb) Save(a Account) (*Account, error) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?, ?, ?, ?, ?)"

	res, err := d.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting inserted Account ID: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, error) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values(?, ?, ?, ?)`, t.AccountID, t.Amount, t.Type, t.Date)

	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountID)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		logger.Error("Erro while getting the last transaction id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	account, err := d.FindBy(t.AccountID)
	if err != nil {
		return nil, err
	}
	t.ID = strconv.FormatInt(transactionID, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(ID string) (*Account, error) {
	accountSQL := "SELECT amount FROM accounts WHERE account_id = ?"

	var account Account
	err := d.client.Get(&account, accountSQL, ID)
	if err != nil {
		logger.Error("Error while scanning account " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Account not found")
		}
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return &account, nil
}
