package transaction

import (
	"net/http"

	"github.com/hpazk/go-booklib/auth"
	"github.com/hpazk/go-booklib/helper"
	"github.com/labstack/echo/v4"
)

type handler struct {
	services Services
	auth     auth.AuthServices
}

func transactionHandler(services Services, authServices auth.AuthServices) *handler {
	return &handler{services, authServices}
}

func (h *handler) PostTransaction(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "post-transaction"})
}

func (h *handler) GetTransactions(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "get-transactions"})
}

func (h *handler) GetTransaction(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "get-transaction"})
}

func (h *handler) PutTransaction(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "put-transaction"})
}

func (h *handler) DeleteTransaction(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "delete-transaction"})
}