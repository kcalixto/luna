package types

type IncomePayload struct {
	Amount float64 `json:"amount" validate:"required"`
}

type ExpensePayload struct {
	Amount float64 `json:"amount" validate:"required"`
}
