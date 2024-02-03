package handlers

import (
	"github.com/labstack/echo/v4"
)

func Response(c echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, data)
}
