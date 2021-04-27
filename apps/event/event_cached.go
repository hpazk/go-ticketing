package event

import (
	"encoding/json"
	"net/http"

	"github.com/hpazk/go-ticketing/cache"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetEventsCached(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		events, _ := h.services.FetchEvents()
		response := events
		rd := cache.GetRedisInstance()
		eventsJson, _ := rd.Get("get-events").Result()
		if err := json.Unmarshal([]byte(eventsJson), &events); err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, response)
	}
}
