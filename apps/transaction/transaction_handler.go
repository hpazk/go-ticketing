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
	return c.JSON(http.StatusOK, helper.M{"message": "post-transaction"})
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
	fmt.Println(paramEventID)
	eventId, _ = strconv.Atoi(paramEventID)
	tsxs, _ := h.services.FetchTransactionsByEvent(uint(eventId))
	fmt.Println(tsxs)

	var transactions []Transaction
	rd := cache.GetRedisInstance()
	tsxsJson, _ := rd.Get("tsx").Result()
	if err := json.Unmarshal([]byte(tsxsJson), &transactions); err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, transactions)
}
