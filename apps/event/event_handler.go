package event

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

func eventHandler(services Services, authServices auth.AuthServices) *handler {
	return &handler{services, authServices}
}

func (h *handler) PostEvent(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "post-event"})
}

func (h *handler) GetEvents(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "get-events"})
}

func (h *handler) GetEvent(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "get-event"})
}

func (h *handler) PutEvent(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "put-event"})
}

func (h *handler) DeleteEvent(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "delete-event"})
}
