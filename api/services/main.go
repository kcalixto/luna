package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	"github.com/kcalixto/mojo-jojo/api/services/finances"
)

type Services struct {
	Finances financesService.FinancesService
}

func New(
	vld *validator.Validate,
	cfg *config.Config,
) *Services {
	return &Services{
		Finances: *financesService.NewFinancesService(vld, cfg),
	}
}
