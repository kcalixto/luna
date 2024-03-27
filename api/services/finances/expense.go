package financesService

import "github.com/kcalixto/mojo-jojo/api/types"

type FinancesExpense struct {
	svc *FinancesService
}

func newFinancesExpenseService(svc *FinancesService) IFinancesExpense {
	return &FinancesExpense{svc}
}

func (f *FinancesExpense) Add(request types.ExpensePayload) (response string, err error) {
	// TODO
	return "expense added successfully", nil
}

func (f *FinancesExpense) Put(request types.ExpensePayload) (response string, err error) {
	// TODO
	return "expense updated successfully", nil
}

func (f *FinancesExpense) Delete(request types.ExpensePayload) (response string, err error) {
	// TODO
	return "expense deleted successfully", nil
}
