package ports

import (
	"context"

	"github.com/gofrs/uuid"
)

type AccountRepository interface {
	OpenTransaction() (err error, cancelContext context.CancelFunc)
	Commit() error
	Rollback() error
	SelectBalanceByUserID(userID uuid.UUID) (int, error)
	RemoveBalanceByUserID(amount int, userID uuid.UUID) error
	AddBalanceByUserID(amount int, userID uuid.UUID) error
}
