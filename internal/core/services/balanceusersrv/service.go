package balanceusersrv

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/leonardodelira/go-api-transfer-money/internal/core/ports"
)

type service struct {
	accountRepo ports.AccountRepository
}

func New(accountRepo ports.AccountRepository) ports.BalanceService {
	return &service{
		accountRepo: accountRepo,
	}
}

func (s *service) GetBalance(ctx context.Context, userID uuid.UUID) (int, error) {
	s.accountRepo.OpenTransaction()
	balance, err := s.accountRepo.SelectBalanceByUserID(userID)
	s.accountRepo.Commit()
	if err != nil {
		return 0, err
	}
	return balance, nil
}
