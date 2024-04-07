package controllerutils

import (
	"time"

	"github.com/kcalixto/mojo-jojo/api/controllers/viewmodels"
	"github.com/kcalixto/mojo-jojo/api/types"
)

const (
	DEFAULT_DATE_FORMAT = time.RFC3339
)

func ParseIncomePayload(request viewmodels.IncomeRequestPayload) (payload types.Income, err error) {
	date, err := time.Parse(DEFAULT_DATE_FORMAT, request.Date)
	if err != nil {
		return types.Income{}, err
	}

	return types.Income{
		Date:        date,
		Amount:      request.Amount,
		Received:    *request.Received,
		Description: request.Description,
		Category:    request.Category,
		Account:     request.Account,
		Recurrent:   *request.Recurrent,
		Note:        request.Note,
		Ignore:      *request.Ignore,
	}, nil
}

func ParseExpensePayload(request viewmodels.ExpenseRequestPayload) (payload types.Expense, err error) {
	date, err := time.Parse(DEFAULT_DATE_FORMAT, request.Date)
	if err != nil {
		return types.Expense{}, err
	}

	return types.Expense{
		Amount:      request.Amount,
		Received:    *request.Received,
		Date:        date,
		Description: request.Description,
		Category:    request.Category,
		Account:     request.Account,
		Recurrent:   *request.Recurrent,
		Note:        request.Note,
		Ignore:      *request.Ignore,
	}, nil
}
