package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leonardodelira/go-api-transfer-money/dependencies"
)

func Start() {
	serverURL := fmt.Sprintf(":%d", 3000) //todo: env

	server := echo.New()
	server.Use(middleware.Logger())

	dependencies.Init()
	Routes(server)
	server.Logger.Fatal(server.Start(serverURL))
}
