package financesService

import (
	"context"

	financesRepository "github.com/kcalixto/luna/api/data/repository/finances"
	"github.com/kcalixto/luna/api/types"
)

type FinancesExpenseService struct {
	svc  *FinancesService
	repo financesRepository.IFinancesExpenseRepository
}

func newFinancesExpenseService(svc *FinancesService, repoManager *financesRepository.FinancesRepositoryManager) (IFinancesExpenseService, error) {
	repo, err := repoManager.NewFinancesExpenseRepository()
	if err != nil {
		return nil, err
	}

	return &FinancesExpenseService{svc, repo}, nil
}

func (s *FinancesExpenseService) Add(ctx context.Context, payload types.Expense) (response string, err error) {
	payload.ID = s.svc.utils.GenerateUUID()

	err = s.repo.AddExpense(ctx, payload)
	if err != nil {
		return response, err
	}

	return "expense added successfully", nil
}

// func (s *FinancesExpenseService) Put(request types.ExpensePayload) (response string, err error) {
// 	// TODO
// 	return "expense updated successfully", nil
// }

// func (s *FinancesExpenseService) Delete(request types.ExpensePayload) (response string, err error) {
// 	// TODO
// 	return "expense deleted successfully", nil
// }
