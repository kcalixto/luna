package financesService

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	"github.com/kcalixto/mojo-jojo/api/data/repository"
	"github.com/kcalixto/mojo-jojo/api/services/serviceUtils"
	"github.com/kcalixto/mojo-jojo/api/types"
)

type FinancesService struct {
	vld     *validator.Validate
	cfg     *config.Config
	repo    *repository.RepositoryManager
	utils   *serviceUtils.ServiceUtils
	Income  IFinancesIncomeService
	Expense IFinancesExpenseService
}

func NewFinancesService(
	vld *validator.Validate,
	cfg *config.Config,
	repo *repository.RepositoryManager,
	utils *serviceUtils.ServiceUtils,
) (svc *FinancesService, err error) {
	svc = &FinancesService{
		vld:   vld,
		cfg:   cfg,
		repo:  repo,
		utils: utils,
	}

	financesRepoManager := repo.NewFinancesRepositoryManager()

	svc.Income, err = newFinancesIncomeService(svc, financesRepoManager)
	if err != nil {
		return nil, err
	}
	svc.Expense = newFinancesExpenseService(svc)

	return svc, nil
}

type IFinancesIncomeService interface {
	Add(context.Context, types.IncomePayload) (string, error)
}
type IFinancesExpenseService interface {
	Add(context.Context, types.ExpensePayload) (string, error)
}
