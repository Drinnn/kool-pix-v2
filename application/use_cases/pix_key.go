package use_cases

import "github.com/Drinnn/kool-pix-v2/domain/models"

type PixKeyUsecase struct {
	PixKeyRepository models.PixKeyRepositoryInterface
}

func (uc *PixKeyUsecase) RegisterKey(key string, kind string, accountID string) (*models.PixKey, error) {
	account, err := uc.PixKeyRepository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := models.NewPixKey(account, kind, key)
	if err != nil {
		return nil, err
	}

	uc.PixKeyRepository.Register(pixKey)

	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (uc *PixKeyUsecase) FindKey(key string, kind string) (*models.PixKey, error) {
	pixKey, err := uc.PixKeyRepository.FindByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
