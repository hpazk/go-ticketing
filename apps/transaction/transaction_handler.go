package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-ticketing/auth"
	"github.com/hpazk/go-ticketing/cache"
	"github.com/hpazk/go-ticketing/database/model"
	"github.com/hpazk/go-ticketing/helper"

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

	var participant model.User
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	participant.ID = uint(claims["user_id"].(float64))
	participant.Email = claims["user_email"].(string)

	newTransaction, err := h.services.SaveTransaction(req, participant)
	if err != nil {
		response := helper.ResponseFormatterWD(http.StatusInternalServerError, "fail", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}
	transactionData := transactionFormatter(newTransaction)
	response := helper.ResponseFormatter(http.StatusOK, "success", "chekout success", transactionData)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetTransactions(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" && role != "creator" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	rd := cache.GetRedisInstance()
	tsxs, _ := h.services.FetchTransactions()

	tsxsFormed := tsxsFormatter(tsxs)
	tsxsJson, _ := json.Marshal(tsxsFormed)

	jsonString := string(tsxsJson)

	rd.Set("tsx", jsonString, time.Hour*1)

	return c.JSON(http.StatusOK, tsxs)
}

func (h *handler) GetTransaction(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" && role != "creator" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	return c.JSON(http.StatusOK, helper.M{"message": "get-transaction"})
}

func (h *handler) PatchTransaction(c echo.Context) error {
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
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", "invalid request")

		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"fail": errorFormatter}
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", errorMessage)

		return c.JSON(http.StatusBadRequest, response)
	}

	err := h.services.EditTransaction(uint(id), req)
	if err != nil {
		response := helper.ResponseFormatterWD(http.StatusBadRequest, "fail", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseFormatterWD(http.StatusOK, "success", "transaction successfully updated")
	return c.JSON(http.StatusOK, response)
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

// Upload Photo Handler
func (h *handler) PostPaymentConfirmation(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	participanID := uint(claims["user_id"].(float64))
	// role := claims["user_role"]

	// Source
	image, err := c.FormFile("image")
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

	imagePath := fmt.Sprintf("public/%d-%s", participanID, image.Filename)

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
	err = h.services.UploadPaymentOrder(participanID, imagePath)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "file upload failed", helper.M{"is_uploaded": false})
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.ResponseFormatter(http.StatusOK, "success", "image succesfully uploaded", helper.M{"is_uploaded": true})

	return c.JSON(http.StatusOK, response)
}
