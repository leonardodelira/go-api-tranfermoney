package transfersrv

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/leonardodelira/go-api-transfer-money/internal/core/ports"
)

type service struct {
	accountRepo ports.AccountRepository
}

func New(accountRepo ports.AccountRepository) ports.TransferService {
	return &service{
		accountRepo: accountRepo,
	}
}

func (s *service) TransferMoney(ctx context.Context, amount int, debtorID, beneficiaryID uuid.UUID) error {
	s.accountRepo.OpenTransaction()
	balanceDebtor, err := s.accountRepo.SelectBalanceByUserID(debtorID)
	if err != nil {
		return err
	}
	if balanceDebtor < amount {
		return errors.New("insufficient balance on debtor account")
	}

	err = s.accountRepo.RemoveBalanceByUserID(amount, debtorID)
	if err != nil {
		s.accountRepo.Rollback()
		return err
	}

	err = s.accountRepo.AddBalanceByUserID(amount, beneficiaryID)
	if err != nil {
		s.accountRepo.Rollback()
		return err
	}
	s.accountRepo.Commit()
	return nil
}
