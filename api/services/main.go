package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	"github.com/kcalixto/mojo-jojo/api/data/repository"
	"github.com/kcalixto/mojo-jojo/api/services/account"
	"github.com/kcalixto/mojo-jojo/api/services/finances"
	"github.com/kcalixto/mojo-jojo/api/services/serviceUtils"
)

type Services struct {
	Finances *financesService.FinancesService
	Account  *accountService.AccountService
}

func New(
	vld *validator.Validate,
	cfg *config.Config,
	repo *repository.RepositoryManager,
) (*Services, error) {
	utils := serviceUtils.NewServiceUtils()

	financesSvc, err := financesService.NewFinancesService(vld, cfg, repo, utils)
	if err != nil {
		return nil, err
	}

	accountSvc, err := accountService.NewAccountService(vld, cfg, repo, utils)
	if err != nil {
		return nil, err
	}

	return &Services{
		Finances: financesSvc,
		Account:  accountSvc,
	}, nil
}
