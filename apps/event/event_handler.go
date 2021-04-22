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

	newEvent, err := h.services.SaveEvent(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	eventData := newEvent
	return c.JSON(http.StatusOK, eventData)
}

func (h *handler) GetEvents(c echo.Context) error {
	events, _ := h.services.FetchEvents()
	response := events
	return c.JSON(http.StatusOK, response)
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
