package use_cases

import (
	"errors"

	"github.com/Drinnn/kool-pix-v2/domain/models"
)

type TransactionUseCase struct {
	TransactionRepository models.TransactionRepositoryInterface
	PixKeyRepository      models.PixKeyRepositoryInterface
}

func (uc *TransactionUseCase) Register(accountID string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*models.Transaction, error) {
	account, err := uc.PixKeyRepository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := uc.PixKeyRepository.FindByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := models.NewTransaction(account, pixKey, amount, description)
	if err != nil {
		return nil, err
	}

	uc.TransactionRepository.Save(transaction)
	if transaction.ID == "" {
		return nil, err
	}

	return transaction, errors.New("unable to process transaction")

}

func (uc *TransactionUseCase) Complete(transactionID string) (*models.Transaction, error) {
	transaction, err := uc.TransactionRepository.Find(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Status = models.TRANSACTION_COMPLETED
	err = uc.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Error(transactionID string) (*models.Transaction, error) {
	transaction, err := uc.TransactionRepository.Find(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Status = models.TRANSACTION_ERROR
	err = uc.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
