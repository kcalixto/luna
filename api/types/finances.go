package types

type IncomePayload struct {
	ID     string  `json:"id"`
	Title  string  `json:"title" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type ExpensePayload struct {
	ID     string  `json:"id"`
	Title  string  `json:"title" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}
