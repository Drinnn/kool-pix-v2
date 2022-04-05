package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TRANSACTION_PENDING   string = "pending"
	TRANSACTION_COMPLETED string = "completed"
	TRANSACTION_ERROR     string = "error"
	TRANSACTION_CONFIRMED string = "confirmed"
)

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
}

type Transactions struct {
	Transactions []Transaction
}

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if transaction.PixKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, pixKeyTo *PixKey, amount float64, description string) (*Transaction, error) {
	transaction := &Transaction{
		AccountFrom: accountFrom,
		PixKeyTo:    pixKeyTo,
		Amount:      amount,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.Status = TRANSACTION_PENDING
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TRANSACTION_COMPLETED
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TRANSACTION_CONFIRMED
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TRANSACTION_ERROR
	transaction.CancelDescription = description
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return err
	}

	return nil
}
