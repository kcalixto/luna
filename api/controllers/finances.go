package controllers

import (
	"context"
	controllerutils "github.com/kcalixto/luna/api/controllers/controllerUtils"
	"github.com/kcalixto/luna/api/controllers/viewmodels"
	"github.com/kcalixto/luna/api/services"
)

type FinancesController struct {
	svc *services.Services
}

func newFinancesController(svc *services.Services) *FinancesController {
	return &FinancesController{svc}
}

func (f *FinancesController) HandleIncome(ctx context.Context, requestUri viewmodels.IncomeRequestUri, request viewmodels.IncomeRequestPayload) (response string, err error) {
	payload, err := controllerutils.ParseIncomePayload(request)
	if err != nil {
		return response, err
	}

	switch requestUri.Type {
	case viewmodels.IncomeActionEnumAdd:
		response, err = f.svc.Finances.Income.Add(ctx, payload)
		if err != nil {
			return response, err
		}

		return response, nil
	case viewmodels.IncomeActionEnumPut:
		return "WIP", nil
	case viewmodels.IncomeActionEnumDelete:
		return "WIP", nil
	}

	return "action not found", nil
}

func (f *FinancesController) HandleExpense(ctx context.Context, requestUri viewmodels.ExpenseRequestUri, request viewmodels.ExpenseRequestPayload) (response string, err error) {
	payload, err := controllerutils.ParseExpensePayload(request)
	if err != nil {
		return response, err
	}

	switch requestUri.Type {
	case viewmodels.IncomeActionEnumAdd:
		response, err = f.svc.Finances.Expense.Add(ctx, payload)
		if err != nil {
			return response, err
		}

		return response, nil
	case viewmodels.IncomeActionEnumPut:
		return "WIP", nil
	case viewmodels.IncomeActionEnumDelete:
		return "WIP", nil
	}

	return response, nil
}
