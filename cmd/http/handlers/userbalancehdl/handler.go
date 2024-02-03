package userbalancehdl

import (
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leonardodelira/go-api-transfer-money/cmd/http/handlers"
	"github.com/leonardodelira/go-api-transfer-money/internal/core/ports"
)

type handlerhttp struct {
	service ports.BalanceService
}

func NewHttpHandler(service ports.BalanceService) handlerhttp {
	return handlerhttp{
		service: service,
	}
}

func (h *handlerhttp) GetBalanceUserHandler(c echo.Context) error {
	userID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return errors.New("invalid userID") //todo: improve handler error
	}
	balance, err := h.service.GetBalance(c.Request().Context(), userID)
	if err != nil {
		return err
	}
	r := &Response{
		Balance: balance,
	}
	return handlers.Response(c, http.StatusOK, r)
}
