package financesRepository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	financesModels "github.com/kcalixto/mojo-jojo/api/data/models/finances"
	"github.com/kcalixto/mojo-jojo/api/types"
)

// IFinancesExpenseRepository is the interface that defines the methods for the finances expense repository
type IFinancesExpenseRepository interface {
	AddExpense(ctx context.Context, income types.Expense) (err error)
	PutExpense(ctx context.Context) (err error)
	DeleteExpense(ctx context.Context) (err error)
}

// financesExpenseRepository is the struct that implements the IFinancesExpenseRepository interface and handle the database operations for the finances expense
type financesExpenseRepository struct {
	table  string
	client *dynamodb.Client
	vld    *validator.Validate
}

// creates a new instance of IFianancesExpenseRepository
func newFinancesExpensesRepository(cfg *config.Database, client *dynamodb.Client, vld *validator.Validate) (IFinancesExpenseRepository, error) {
	table := cfg.SingleTableName
	if table == "" {
		return nil, fmt.Errorf("newFinancesIncomesRepository -> table name is empty")
	}

	return &financesExpenseRepository{table, client, vld}, nil
}

// AddIncome adds a new income to the database
func (r *financesExpenseRepository) AddExpense(ctx context.Context, expense types.Expense) (err error) {
	item := financesModels.AddIncomeModel{
		PK:          r.buildPK(),
		SK:          r.buildSK(expense.ID),
		Amount:      expense.Amount,
		Received:    expense.Received,
		Date:        expense.Date.Format("2006-01-02"),
		Description: expense.Description,
		Category:    expense.Category,
		Account:     expense.Account,
		Recurrent:   expense.Recurrent,
		Note:        expense.Note,
		Ignore:      expense.Ignore,
	}

	marshaledItem, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      marshaledItem,
	})
	if err != nil {
		message := fmt.Sprintf("AddExpense -> got error calling PutItem: %s", err.Error())
		return fmt.Errorf(message)
	}

	return nil
}

// PutExpense updates an existing expense in the database
func (r *financesExpenseRepository) PutExpense(ctx context.Context) (err error) {
	return nil
}

// DeleteExpense deletes an existing expense from the database
func (r *financesExpenseRepository) DeleteExpense(ctx context.Context) (err error) {
	return nil
}

// Builds the partition key for the income inside the single table database schema
func (r *financesExpenseRepository) buildPK() string {
	return "ACCOUNT#DEFAULT"
}

// Builds the sort key for the income inside the single table database schema
func (r *financesExpenseRepository) buildSK(id string) string {
	return fmt.Sprintf("EXPENSE#%s", id)
}
