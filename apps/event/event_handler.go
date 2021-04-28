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
		response := helper.ResponseFormatterWD(http.StatusUnauthorized, "fail", "Please provide valid credentials")
		return c.JSON(http.StatusUnauthorized, response)
	}
	req := new(request)

	// Check request
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", "invalid request")

		return c.JSON(http.StatusBadRequest, response)
	}

	// Validate request
	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", errorMessage)

		return c.JSON(http.StatusBadRequest, response)
	}

	err := h.services.SaveEvent(req)
	if err != nil {
		response := helper.ResponseFormatterWD(http.StatusInternalServerError, "fail", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.ResponseFormatterWD(http.StatusOK, "success", "event successfully drafted")
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetEvents(c echo.Context) error {
	events, err := h.services.FetchEvents()
	if err != nil {
		response := helper.ResponseFormatterWD(http.StatusInternalServerError, "fail", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := eventsResponseFormatter(events)
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
		response := helper.ResponseFormatterWD(http.StatusUnauthorized, "fail", "Please provide valid credentials")
		return c.JSON(http.StatusUnauthorized, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	req := new(updateRequest)

	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", "invalid request")

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"fail": errorFormatter}
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", errorMessage)

		return c.JSON(http.StatusBadRequest, response)
	}

	err := h.services.EditEvent(uint(id), req)
	if err != nil {
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseFormatterWD(http.StatusOK, "success", "event successfully updated")
	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteEvent(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
		response := helper.ResponseFormatterWD(http.StatusUnauthorized, "fail", "Please provide valid credentials")
		return c.JSON(http.StatusUnauthorized, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.services.RemoveEvent(uint(id)); err != nil {
		response := helper.ResponseFormatterWD(http.StatusInternalServerError, "fail", err)
		return c.JSON(http.StatusOK, response)
	}

	message := fmt.Sprintf("event %d was deleted", id)
	response := helper.ResponseFormatterWD(http.StatusOK, "success", message)
	return c.JSON(http.StatusOK, response)
}
