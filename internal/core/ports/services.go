package ports

import (
	"context"

	"github.com/gofrs/uuid"
)

type TransferService interface {
	TransferMoney(ctx context.Context, amount int, debtorID, beneficiaryID uuid.UUID) error
}

type BalanceService interface {
	GetBalance(ctx context.Context, userID uuid.UUID) (int, error)
}
