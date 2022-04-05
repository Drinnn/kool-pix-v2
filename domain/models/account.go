package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string `json:"owner_name" valid:"required"`
	Bank      *Bank  `valid:"-"`
	Number    string `json:"number" valid:"notnull"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, ownerName string, number string) (*Account, error) {
	account := &Account{
		Bank:      bank,
		OwnerName: ownerName,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return account, nil
}