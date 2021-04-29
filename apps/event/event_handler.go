package event

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

func (h *handler) PatchEventBanner(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	// participanID := uint(claims["user_id"].(float64))
	role := claims["user_role"]

	if role != "admin" && role != "creator" {
		response := helper.ResponseFormatterWD(http.StatusUnauthorized, "fail", "Please provide valid credentials")
		return c.JSON(http.StatusUnauthorized, response)
	}

	paramEventID := c.Param("id")
	eventID, _ := strconv.Atoi(paramEventID)

	// Source
	image, err := c.FormFile("banner")
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusBadRequest, response)
	}

	src, err := image.Open()
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer src.Close()

	imagePath := fmt.Sprintf("public/banner/%d-%s", eventID, image.Filename)

	// Destination
	dst, err := os.Create(imagePath)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Upload
	err = h.services.UploadBanner(uint(eventID), imagePath)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.ResponseFormatter(http.StatusOK, "success", "image succesfully uploaded", helper.M{"is_uploaded": true})

	return c.JSON(http.StatusOK, response)
}
