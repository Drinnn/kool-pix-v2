package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(account *Account, kind string, key string) (*PixKey, error) {
	pixKey := &PixKey{
		Account:   account,
		AccountID: account.ID,
		Kind:      kind,
		Key:       key,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.Status = "active"
	pixKey.CreatedAt = time.Now()
	pixKey.UpdatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
