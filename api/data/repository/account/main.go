package accountRepository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/luna/api/config"
)

type AccountRepositoryManager struct {
	client *dynamodb.Client
	cfg    *config.Database
	vld    *validator.Validate
}

func NewAccountRepositoryManager(client *dynamodb.Client, cfg *config.Database, vld *validator.Validate) *AccountRepositoryManager {
	return &AccountRepositoryManager{client, cfg, vld}
}

func (m *AccountRepositoryManager) NewAccountTransactionRepository() (IAccountTransactionRepository, error) {
	return newAccountTransactionRepository(m.cfg, m.client, m.vld)
}
