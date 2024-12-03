package accountService

import (
	"context"

	"github.com/kcalixto/luna/api/data/repository/account"
	"github.com/kcalixto/luna/api/types"
)

type AccountTransactionService struct {
	svc  *AccountService
	repo accountRepository.IAccountTransactionRepository
}

func newAccountTransactionService(svc *AccountService, repoManager *accountRepository.AccountRepositoryManager) (IAccountTransactionService, error) {
	repo, err := repoManager.NewAccountTransactionRepository()
	if err != nil {
		return nil, err
	}

	return &AccountTransactionService{svc, repo}, nil
}

func (s *AccountTransactionService) List(ctx context.Context) (response []types.Transaction, err error) {
	transactions, err := s.repo.List(ctx)
	if err != nil {
		return response, err
	}

	for _, transaction := range transactions {
		t, err := transaction.ToType()
		if err != nil {
			return response, err
		}

		response = append(response, t)
	}

	return response, nil
}
