package dependencies

import (
	"github.com/leonardodelira/go-api-transfer-money/internal/core/services/balanceusersrv"
	transfersrv "github.com/leonardodelira/go-api-transfer-money/internal/core/services/transfer"
)

func initServices() {
	BalanceService = balanceusersrv.New(AccountRepository)
	TransferService = transfersrv.New(AccountRepository)
}
