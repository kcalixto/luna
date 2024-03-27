package financesModels

type AddIncomeModel struct {
	PK     string  `validate:"required" dynamodbav:"pk"`
	SK     string  `validate:"required" dynamodbav:"sk"`
	Title  string  `validate:"required" dynamodbav:"title"`
	Amount float64 `validate:"required" dynamodbav:"amount"`
}
