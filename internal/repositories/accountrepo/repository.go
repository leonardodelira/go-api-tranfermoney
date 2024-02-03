package accountrepo

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leonardodelira/go-api-transfer-money/internal/core/ports"
)

const transactionTimeout = 20 * time.Second
const defaultTimeout = 10 * time.Second

type postgresRepo struct {
	Conn *pgxpool.Pool
	tx   pgx.Tx
	ctx  context.Context
}

func NewAccountRepoPostgres(conn *pgxpool.Pool) ports.AccountRepository {
	return &postgresRepo{
		Conn: conn,
	}
}

func (repo *postgresRepo) OpenTransaction() (err error, cancelContext context.CancelFunc) {
	ctx, cancelContext := context.WithTimeout(context.Background(), transactionTimeout)

	repo.ctx = ctx
	options := pgx.TxOptions{}
	repo.tx, err = repo.Conn.BeginTx(ctx, options)

	return
}

func (repo *postgresRepo) Rollback() error {
	if repo.tx.Conn().IsClosed() {
		return errors.New("database transaction is closed already")
	}

	return repo.tx.Rollback(repo.ctx)
}

func (repo *postgresRepo) Commit() error {
	if repo.tx.Conn().IsClosed() {
		return errors.New("database transaction is closed already")
	}

	return repo.tx.Commit(repo.ctx)
}

func (repo *postgresRepo) SelectBalanceByUserID(userID uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(repo.ctx, defaultTimeout)
	defer cancel()

	sql := "SELECT amount FROM accounts WHERE user_id = $1"

	row := repo.tx.QueryRow(ctx, sql, userID)
	amount := 0

	if err := row.Scan(
		&amount,
	); err != nil {
		return 0, err
	}

	return amount, nil
}

func (repo *postgresRepo) RemoveBalanceByUserID(amount int, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(repo.ctx, defaultTimeout)
	defer cancel()

	sql := "UPDATE accounts SET amount = amount - $1 WHERE user_id = $2"

	_, err := repo.tx.Exec(ctx, sql, amount, userID)

	return err
}

func (repo *postgresRepo) AddBalanceByUserID(amount int, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(repo.ctx, defaultTimeout)
	defer cancel()

	sql := "UPDATE accounts SET amount = amount + $1 WHERE user_id = $2"

	_, err := repo.tx.Exec(ctx, sql, amount, userID)

	return err
}
