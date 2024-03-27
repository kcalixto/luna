package financesRepository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
)

type FinancesRepositoryManager struct {
	client *dynamodb.Client
	cfg    *config.Database
	vld    *validator.Validate
}

func NewFinancesRepositoryManager(client *dynamodb.Client, cfg *config.Database, vld *validator.Validate) *FinancesRepositoryManager {
	return &FinancesRepositoryManager{client, cfg, vld}
}

func (m *FinancesRepositoryManager) NewFinancesIncomeRepository() (IFinancesIncomeRepository, error) {
	return newFinancesIncomesRepository(m.cfg, m.client, m.vld)
}
