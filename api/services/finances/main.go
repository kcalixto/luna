package financesService

import (
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	"github.com/kcalixto/mojo-jojo/api/types"
)

type FinancesService struct {
	vld     *validator.Validate
	cfg     *config.Config
	Income  IFinancesIncome
	Expense IFinancesExpense
}

func NewFinancesService(
	vld *validator.Validate,
	cfg *config.Config,
) *FinancesService {
	svc := &FinancesService{
		vld: vld,
		cfg: cfg,
	}

	svc.Income = newFinancesIncomeService(svc)
	svc.Expense = newFinancesExpenseService(svc)

	return svc
}

type IFinancesIncome interface {
	Add(request types.IncomePayload) (response string, err error)
}
type IFinancesExpense interface {
	Add(request types.ExpensePayload) (response string, err error)
}
