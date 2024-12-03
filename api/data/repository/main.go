package repository

import (
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/luna/api/config"
	accountRepository "github.com/kcalixto/luna/api/data/repository/account"
	financesRepository "github.com/kcalixto/luna/api/data/repository/finances"
)

type RepositoryManager struct {
	client *dynamodb.Client
	cfg    *config.Database
	vld    *validator.Validate
}

func New(ctx context.Context, cfg *config.Config, vld *validator.Validate) *RepositoryManager {
	// Load AWS credentials and config
	awscfg, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("unable to load SDK config, ", err)
	}

	// Create DynamoDB client
	client := dynamodb.NewFromConfig(awscfg)

	return &RepositoryManager{client, &cfg.Database, vld}
}

func (m *RepositoryManager) NewFinancesRepositoryManager() *financesRepository.FinancesRepositoryManager {
	return financesRepository.NewFinancesRepositoryManager(m.client, m.cfg, m.vld)
}

func (m *RepositoryManager) NewAccountRepositoryManager() *accountRepository.AccountRepositoryManager {
	return accountRepository.NewAccountRepositoryManager(m.client, m.cfg, m.vld)
}
