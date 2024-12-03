package financesRepository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/luna/api/config"
	financesModels "github.com/kcalixto/luna/api/data/models/finances"
	"github.com/kcalixto/luna/api/types"
)

// IFinancesIncomeRepository is the interface that defines the methods for the finances income repository
type IFinancesIncomeRepository interface {
	AddIncome(ctx context.Context, income types.Income) (err error)
	PutIncome(ctx context.Context) (err error)
	DeleteIncome(ctx context.Context) (err error)
}

// financesIncomeRepository is the struct that implements the IFinancesIncomeRepository interface and handle the database operations for the finances income
type financesIncomeRepository struct {
	table  string
	client *dynamodb.Client
	vld    *validator.Validate
}

// creates a new instance of IFianancesIncomeRepository
func newFinancesIncomesRepository(cfg *config.Database, client *dynamodb.Client, vld *validator.Validate) (IFinancesIncomeRepository, error) {
	table := cfg.SingleTableName
	if table == "" {
		return nil, fmt.Errorf("newFinancesIncomesRepository -> table name is empty")
	}

	return &financesIncomeRepository{table, client, vld}, nil
}

// AddIncome adds a new income to the database
func (r *financesIncomeRepository) AddIncome(ctx context.Context, income types.Income) (err error) {
	item := financesModels.AddIncomeModel{
		PK:          r.buildPK(),
		SK:          r.buildSK(income.ID),
		Amount:      income.Amount,
		Received:    income.Received,
		Date:        income.Date.Format("2006-01-02"),
		Description: income.Description,
		Category:    income.Category,
		Account:     income.Account,
		Recurrent:   income.Recurrent,
		Note:        income.Note,
		Ignore:      income.Ignore,
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
		message := fmt.Sprintf("AddIncome -> got error calling PutItem: %s", err.Error())
		return fmt.Errorf(message)
	}

	return nil
}

// PutIncome updates an existing income in the database
func (r *financesIncomeRepository) PutIncome(ctx context.Context) (err error) {
	return nil
}

// DeleteIncome deletes an existing income from the database
func (r *financesIncomeRepository) DeleteIncome(ctx context.Context) (err error) {
	return nil
}

// Builds the partition key for the income inside the single table database schema
func (r *financesIncomeRepository) buildPK() string {
	return "ACCOUNT#DEFAULT"
}

// Builds the sort key for the income inside the single table database schema
func (r *financesIncomeRepository) buildSK(id string) string {
	return fmt.Sprintf("INCOME#%s", id)
}
