package transferhdl

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leonardodelira/go-api-transfer-money/cmd/http/handlers"
	"github.com/leonardodelira/go-api-transfer-money/internal/core/ports"
)

type transferRequest struct {
	Amount        int       `json:"amount"`
	DebtorID      uuid.UUID `json:"debtorId"`
	BeneficiaryID uuid.UUID `json:"beneficiaryId"`
}

type handlerhttp struct {
	service ports.TransferService
}

func NewHttpHandler(service ports.TransferService) handlerhttp {
	return handlerhttp{
		service: service,
	}
}

func (h *handlerhttp) ExecuteTransferMoneyHandler(c echo.Context) error {
	transferRequest := &transferRequest{}
	if err := c.Bind(&transferRequest); err != nil {
		msgError := "error on bind request params"
		fmt.Print(msgError, err)
		return errors.New(msgError)
	}
	err := h.service.TransferMoney(c.Request().Context(), transferRequest.Amount, transferRequest.DebtorID, transferRequest.BeneficiaryID)
	if err != nil {
		return errors.New("error on transfer money")
	}
	return handlers.Response(c, http.StatusOK, nil)
}
