package types

import "time"

// Transaction is the payload for list transactions inside the api
// for now i'm trying to discover where to put this and how to exactly handle
// incomes/expenses/transactions/etc
type Transaction struct {
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
