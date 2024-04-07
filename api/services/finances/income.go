package financesService

import (
	"context"

	financesRepository "github.com/kcalixto/mojo-jojo/api/data/repository/finances"
	"github.com/kcalixto/mojo-jojo/api/types"
)

type FinancesIncomeService struct {
	svc  *FinancesService
	repo financesRepository.IFinancesIncomeRepository
}

func newFinancesIncomeService(svc *FinancesService, repoManager *financesRepository.FinancesRepositoryManager) (IFinancesIncomeService, error) {
	repo, err := repoManager.NewFinancesIncomeRepository()
	if err != nil {
		return nil, err
	}

	return &FinancesIncomeService{svc, repo}, nil
}

func (s *FinancesIncomeService) Add(ctx context.Context, payload types.Income) (response string, err error) {
	payload.ID = s.svc.utils.GenerateUUID()

	err = s.repo.AddIncome(ctx, payload)
	if err != nil {
		return response, err
	}

	return "income added successfully", nil
}

// func (s *FinancesIncomeService) Put(request types.IncomePayload) (response string, err error) {
// 	// TODO
// 	return "income updated successfully", nil
// }

// func (s *FinancesIncomeService) Delete(request types.IncomePayload) (response string, err error) {
// 	// TODO
// 	return "income deleted successfully", nil
// }
