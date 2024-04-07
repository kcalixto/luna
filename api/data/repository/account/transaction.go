package accountRepository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	"github.com/kcalixto/mojo-jojo/api/data/models/account"
)

// IAccountTransactionRepository is the interface that defines the methods for the account transaction repository
type IAccountTransactionRepository interface {
	List(ctx context.Context) (transactions []models.TransactionModel, err error)
}

// accountTransactionRepository is the struct that implements the IAccountTransactionRepository interface and handle the database operations for the account transaction
type accountTransactionRepository struct {
	table  string
	client *dynamodb.Client
	vld    *validator.Validate
}

// creates a new instance of IAccountTransactionRepository
func newAccountTransactionRepository(cfg *config.Database, client *dynamodb.Client, vld *validator.Validate) (IAccountTransactionRepository, error) {
	table := cfg.SingleTableName
	if table == "" {
		return nil, fmt.Errorf("newAccountTransactionRepository -> table name is empty")
	}

	return &accountTransactionRepository{table, client, vld}, nil
}

// List lists the account transactions from the database
func (r *accountTransactionRepository) List(ctx context.Context) (transactions []models.TransactionModel, err error) {
	keyCond := expression.Key("pk").Equal(expression.Value(r.buildPK()))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		return transactions, err
	}

	queryPaginator := dynamodb.NewQueryPaginator(r.client, &dynamodb.QueryInput{
		TableName:                 aws.String(r.table),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	for queryPaginator.HasMorePages() {
		response, err := queryPaginator.NextPage(ctx)
		if err != nil {
			return transactions, err
		} else {
			var transaction []models.TransactionModel
			err = attributevalue.UnmarshalListOfMaps(response.Items, &transaction)
			if err != nil {
				return transactions, err
			} else {
				transactions = append(transactions, transaction...)
			}
		}
	}

	return transactions, nil
}

// Builds the partition key for the income inside the single table database schema
func (r *accountTransactionRepository) buildPK() string {
	return "ACCOUNT#DEFAULT"
}
