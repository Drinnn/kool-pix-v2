package repository

import (
	"fmt"

	"github.com/Drinnn/kool-pix-v2/domain/models"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r *PixKeyRepositoryDb) Register(pixKey *models.PixKey) (*models.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (r *PixKeyRepositoryDb) FindByKind(key string, kind string) (*models.PixKey, error) {
	var pixKey models.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

func (r *PixKeyRepositoryDb) AddBank(bank *models.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PixKeyRepositoryDb) AddAccount(account *models.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PixKeyRepositoryDb) FindAccount(id string) (*models.Account, error) {
	var account models.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}

	return &account, nil
}

func (r *PixKeyRepositoryDb) FindBank(id string) (*models.Bank, error) {
	var bank models.Bank

	r.Db.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}

	return &bank, nil
}
