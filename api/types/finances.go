package types

import "time"

// Income is the payload for incomes inside the api
type Income struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount" validate:"required"`
	Received    bool      `json:"received" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	Description string    `json:"title" validate:"required"`
	Category    string    `json:"category" validate:"required"`
	Account     string    `json:"account" validate:"required"`
	Recurrent   bool      `json:"recurrent" validate:"required"`
	Note        string    `json:"note" validate:"required"`
	Ignore      bool      `json:"ignore" validate:"required"`
}

// Expense is the payload for expenses inside the api
type Expense struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount" validate:"required"`
	Received    bool      `json:"received" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	Description string    `json:"title" validate:"required"`
	Category    string    `json:"category" validate:"required"`
	Account     string    `json:"account" validate:"required"`
	Recurrent   bool      `json:"recurrent" validate:"required"`
	Note        string    `json:"note" validate:"required"`
	Ignore      bool      `json:"ignore" validate:"required"`
}
