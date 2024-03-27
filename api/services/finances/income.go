package financesService

import "github.com/kcalixto/mojo-jojo/api/types"

type FinancesIncome struct {
	svc *FinancesService
}

func newFinancesIncomeService(svc *FinancesService) IFinancesIncome {
	return &FinancesIncome{svc}
}

func (f *FinancesIncome) Add(request types.IncomePayload) (response string, err error) {
	// TODO
	return "income added successfully", nil
}

func (f *FinancesIncome) Put(request types.IncomePayload) (response string, err error) {
	// TODO
	return "income updated successfully", nil
}

func (f *FinancesIncome) Delete(request types.IncomePayload) (response string, err error) {
	// TODO
	return "income deleted successfully", nil
}
