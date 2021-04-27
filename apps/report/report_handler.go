package report

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-ticketing/auth"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
)

type handler struct {
	services Services
	auth     auth.AuthServices
}

func reportHandler(services Services, authServices auth.AuthServices) *handler {
	return &handler{services, authServices}
}

func (h *handler) GetEventParticipant(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	creatorID := uint(claims["user_id"].(float64))
	role := claims["user_role"]

	if role != "creator" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	statusPayment := c.QueryParam("status_payment")

	report, _ := h.services.FetchEventParticipant(creatorID, statusPayment)
	return c.JSON(http.StatusOK, report)
}
