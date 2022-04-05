package repository

import (
	"fmt"

	"github.com/Drinnn/kool-pix-v2/domain/models"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r *TransactionRepositoryDb) Register(transaction *models.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDb) Save(transaction *models.Transaction) error {
	err := r.Db.Save(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDb) Find(id string) (*models.Transaction, error) {
	var transaction models.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
