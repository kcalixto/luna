package services

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

func newFinancesService(
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

type FinancesIncome struct {
	svc *FinancesService
}

type FinancesExpense struct {
	svc *FinancesService
}

func newFinancesIncomeService(svc *FinancesService) IFinancesIncome {
	return &FinancesIncome{svc}
}

func newFinancesExpenseService(svc *FinancesService) IFinancesExpense {
	return &FinancesExpense{svc}
}

func (f *FinancesIncome) Add(request types.IncomePayload) (response string, err error) {
	// TODO
	return "income added successfully", nil
}

func (f *FinancesExpense) Add(request types.ExpensePayload) (response string, err error) {
	// TODO
	return "expense added successfully", nil
}
