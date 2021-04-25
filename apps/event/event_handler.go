package event

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-ticketing/auth"
	"github.com/hpazk/go-ticketing/helper"
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
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" && role != "creator" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
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
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" && role != "creator" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	req := new(updateRequest)

	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "invalid request", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"fail": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	editedUser, err := h.services.EditEvent(uint(id), req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// TODO updated-formatter
	userData := editedUser

	response := helper.ResponseFormatter(http.StatusOK, "success", "event successfully updated", userData)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteEvent(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.services.RemoveEvent(uint(id)); err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err, nil)
		return c.JSON(http.StatusOK, response)
	}

	message := fmt.Sprintf("event %d was deleted", id)
	response := helper.ResponseFormatter(http.StatusOK, "success", message, nil)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetEventReport(c echo.Context) error {
	// TODO creator-id
	eventID, _ := strconv.Atoi(c.Param("id"))
	report, _ := h.services.FetchEventReport(1, uint(eventID))
	return c.JSON(http.StatusOK, report)
}
