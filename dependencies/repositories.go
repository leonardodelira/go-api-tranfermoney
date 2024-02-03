package dependencies

import (
	"github.com/leonardodelira/go-api-transfer-money/database"
	"github.com/leonardodelira/go-api-transfer-money/internal/repositories/accountrepo"
)

func initRepositories() {
	postgressConn := database.CreateConnectionPostgres()

	AccountRepository = accountrepo.NewAccountRepoPostgres(postgressConn)
}
