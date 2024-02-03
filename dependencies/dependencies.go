package dependencies

import (
	"github.com/leonardodelira/go-api-transfer-money/internal/core/ports"
)

var (
	AccountRepository ports.AccountRepository

	BalanceService  ports.BalanceService
	TransferService ports.TransferService
)

func Init() {
	initRepositories()
	initServices()
}
