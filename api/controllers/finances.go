package controllers

import (
	"github.com/kcalixto/mojo-jojo/api/controllers/viewmodels"
	"github.com/kcalixto/mojo-jojo/api/services"
	"github.com/kcalixto/mojo-jojo/api/types"
)

type FinancesController struct {
	svc *services.Services
}

func newFinancesController(svc *services.Services) *FinancesController {
	return &FinancesController{svc}
}

func (f *FinancesController) HandleIncome(payload viewmodels.IncomePayload) (response string, err error) {
	var p types.IncomePayload
	p.Amount = 100

	response, err = f.svc.Finances.Income.Add(p)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (f *FinancesController) HandleExpense(payload viewmodels.ExpensePayload) (response string, err error) {
	var p types.ExpensePayload
	p.Amount = 100

	response, err = f.svc.Finances.Expense.Add(p)
	if err != nil {
		return response, err
	}

	return response, nil
}
