package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
)

type Services struct {
	Finances FinancesService
}

func New(
	vld *validator.Validate,
	cfg *config.Config,
) *Services {
	return &Services{
		Finances: *newFinancesService(vld, cfg),
	}
}
