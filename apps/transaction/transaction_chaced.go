package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hpazk/go-booklib/cache"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetTransactionsCached(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
}
