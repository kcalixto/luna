package controllers

import (
	"github.com/kcalixto/mojo-jojo/api/services"
)

type Controller struct {
	Finances *FinancesController
}

func New(
	svc *services.Services,
) *Controller {
	return &Controller{
		Finances: newFinancesController(svc),
	}
}
