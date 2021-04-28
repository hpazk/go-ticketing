package event

import (
	"encoding/json"
	"net/http"

	"github.com/hpazk/go-ticketing/cache"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetEventsCached(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := h.services.FetchEvents()
		if err != nil {
			response := helper.ResponseFormatterWD(http.StatusInternalServerError, "fail", err.Error())
			return c.JSON(http.StatusInternalServerError, response)
		}
		response := eventsResponseFormatter(events)
		rd := cache.GetRedisInstance()
		eventsJson, _ := rd.Get("get-events").Result()
		if err := json.Unmarshal([]byte(eventsJson), &events); err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, response)
	}
}
