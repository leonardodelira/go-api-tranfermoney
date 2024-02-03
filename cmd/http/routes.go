package http

import (
	"github.com/labstack/echo/v4"
	"github.com/leonardodelira/go-api-transfer-money/cmd/http/handlers/transferhdl"
	"github.com/leonardodelira/go-api-transfer-money/cmd/http/handlers/userbalancehdl"
	"github.com/leonardodelira/go-api-transfer-money/dependencies"
)

func Routes(server *echo.Echo) {
	userBalanceHandler := userbalancehdl.NewHttpHandler(dependencies.BalanceService)
	transferHandler := transferhdl.NewHttpHandler(dependencies.TransferService)

	server.GET("/users/:id/balance", userBalanceHandler.GetBalanceUserHandler)
	server.POST("/transfer", transferHandler.ExecuteTransferMoneyHandler)
}
