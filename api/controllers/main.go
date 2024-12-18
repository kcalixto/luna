package controllers

import (
	"github.com/kcalixto/luna/api/services"
)

type Controller struct {
	Finances *FinancesController
	Account  *AccountController
}

func New(
	svc *services.Services,
) *Controller {
	return &Controller{
		Finances: newFinancesController(svc),
		Account:  newAccountController(svc),
	}
}
