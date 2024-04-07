package viewmodels

// Income

// IncomeActionEnum is a string enum for the IncomePayload type
type IncomeActionEnum string

// IncomeActionEnum values
// add: add an income
// delete: delete an income
// put: update an income
const (
	IncomeActionEnumAdd    IncomeActionEnum = "add"
	IncomeActionEnumDelete IncomeActionEnum = "delete"
	IncomeActionEnumPut    IncomeActionEnum = "put"
)

// IncomePayload is the payload for the income endpoint
// acton: viewmodel.IncomeActionEnum
// title: title of the income
// amount: amount of the income
type IncomeRequestPayload struct {
	Amount      float64 `json:"amount" binding:"required"`
	Received    *bool   `json:"received" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Account     string  `json:"account" binding:"required"`
	Recurrent   *bool   `json:"recurrent" binding:"required"`
	Note        string  `json:"note" binding:"required"`
	Ignore      *bool   `json:"ignore" binding:"required"`
}
type IncomeRequestUri struct {
	Type IncomeActionEnum `uri:"action" binding:"required"`
}

// Expense

// ExpenseActionEnum is a string enum for the ExpensePayload type
type ExpenseActionEnum string

// ExpenseActionEnum values
// add: add an expense
// delete: delete an expense
// put: update an expense
const (
	ExpenseActionEnumAdd    ExpenseActionEnum = "add"
	ExpenseActionEnumDelete ExpenseActionEnum = "delete"
	ExpenseActionEnumPut    ExpenseActionEnum = "put"
)

// ExpensePayload is the payload for the expense endpoint
// acton: viewmodel.ExpenseActionEnum
// title: title of the expense
// amount: amount of the expense
type ExpenseRequestPayload struct {
	Amount      float64 `json:"amount" binding:"required"`
	Received    *bool   `json:"received" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Account     string  `json:"account" binding:"required"`
	Recurrent   *bool   `json:"recurrent" binding:"required"`
	Note        string  `json:"note" binding:"required"`
	Ignore      *bool   `json:"ignore" binding:"required"`
}
type ExpenseRequestUri struct {
	Type IncomeActionEnum `uri:"action" binding:"required"`
}
