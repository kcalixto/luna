package models

import (
	"time"

	"github.com/kcalixto/luna/api/types"
)

// TransactionModel is the model for getting transactions from the database
type TransactionModel struct {
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

func (m *TransactionModel) ToType() (types.Transaction, error) {
	date, err := time.Parse("2006-01-02", m.Date)
	if err != nil {
		return types.Transaction{}, err
	}

	return types.Transaction{
		ID:          m.PK,
		Amount:      m.Amount,
		Received:    m.Received,
		Date:        date,
		Description: m.Description,
		Category:    m.Category,
		Account:     m.Account,
		Recurrent:   m.Recurrent,
		Note:        m.Note,
		Ignore:      m.Ignore,
	}, nil
}
