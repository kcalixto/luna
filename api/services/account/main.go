package accountService

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/luna/api/config"
	"github.com/kcalixto/luna/api/data/repository"
	"github.com/kcalixto/luna/api/services/serviceUtils"
	"github.com/kcalixto/luna/api/types"
)

type AccountService struct {
	vld         *validator.Validate
	cfg         *config.Config
	repo        *repository.RepositoryManager
	utils       *serviceUtils.ServiceUtils
	Transaction IAccountTransactionService
}

func NewAccountService(
	vld *validator.Validate,
	cfg *config.Config,
	repo *repository.RepositoryManager,
	utils *serviceUtils.ServiceUtils,
) (svc *AccountService, err error) {
	svc = &AccountService{
		vld:   vld,
		cfg:   cfg,
		repo:  repo,
		utils: utils,
	}

	accountRepoManager := repo.NewAccountRepositoryManager()
	if accountRepoManager == nil {
		return nil, errors.New("NewAccountService -> accountRepoManager is nil")
	}

	svc.Transaction, err = newAccountTransactionService(svc, accountRepoManager)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

type IAccountTransactionService interface {
	List(ctx context.Context) (response []types.Transaction, err error)
}
