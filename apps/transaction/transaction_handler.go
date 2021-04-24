package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hpazk/go-booklib/auth"
	"github.com/hpazk/go-booklib/cache"
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
	req := new(request)

	// Check request
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	// Validate request
	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	newTransaction, err := h.services.SaveTransaction(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	transactionData := newTransaction
	return c.JSON(http.StatusOK, transactionData)
}

func (h *handler) GetTransactions(c echo.Context) error {
	rd := cache.GetRedisInstance()
	tsxs, _ := h.services.FetchTransactions()

	tsxsFormed := tsxsFormatter(tsxs)
	tsxsJson, _ := json.Marshal(tsxsFormed)

	jsonString := string(tsxsJson)
	fmt.Println(jsonString)

	rd.Set("tsx", jsonString, time.Hour*1)

	return c.JSON(http.StatusOK, tsxs)
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

func (h *handler) GetTransactionsByEvent(c echo.Context) error {
	var eventId int
	paramEventID := c.Param("id")

	eventId, _ = strconv.Atoi(paramEventID)
	transactions, _ := h.services.FetchTransactionsByEvent(uint(eventId))

	return c.JSON(http.StatusOK, transactions)
}
