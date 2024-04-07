package controllers

import (
	"context"

	"github.com/kcalixto/mojo-jojo/api/services"
	"github.com/kcalixto/mojo-jojo/api/types"
)

type AccountController struct {
	svc *services.Services
}

func newAccountController(svc *services.Services) *AccountController {
	return &AccountController{svc}
}

func (f *AccountController) HandleListTransactions(ctx context.Context) (response []types.Transaction, err error) {
	response, err = f.svc.Account.Transaction.List(ctx)
	if err != nil {
		return response, err
	}

	return response, nil
}
