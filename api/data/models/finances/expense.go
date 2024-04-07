package financesModels

// AddExpenseModel is the model for adding a new expense to the database
type AddExpenseModel struct {
	PK          string  `validate:"required" dynamodbav:"pk"`
	SK          string  `validate:"required" dynamodbav:"sk"`
	Amount      float64 `validate:"required" dynamodbav:"amount"`
	Received    bool    `validate:"required" dynamodbav:"received"`
	Date        string  `validate:"required" dynamodbav:"date"`
	Description string  `validate:"required" dynamodbav:"description"`
	Category    string  `validate:"required" dynamodbav:"category"`
	Account     string  `validate:"required" dynamodbav:"account"`
	Recurrent   bool    `validate:"required" dynamodbav:"recurrent"`
	Note        string  `validate:"required" dynamodbav:"note"`
	Ignore      bool    `validate:"required" dynamodbav:"ignore"`
}
