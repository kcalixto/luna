package financesService

import (
	"context"

	"github.com/kcalixto/mojo-jojo/api/types"
)

type FinancesExpenseService struct {
	svc *FinancesService
}

func newFinancesExpenseService(svc *FinancesService) IFinancesExpenseService {
	return &FinancesExpenseService{svc}
}

func (s *FinancesExpenseService) Add(ctx context.Context, request types.ExpensePayload) (response string, err error) {
	// TODO
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
