package user

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
)

func (h *userHandler) PostNewCreator(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
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

	// Check email exist
	if existEmail := h.userServices.CheckExistEmail(req.Email); existEmail {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", "email is already registered", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// SignUp service
	newUser, err := h.userServices.NewCreator(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "email is already registered", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Get access token
	authToken, err := h.authServices.GetAccessToken(newUser.ID, newUser.Role, newUser.Email)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "something went wrong", nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	// Format data
	userData := userResponseFormatter(newUser, authToken)

	// Passed response
	response := helper.ResponseFormatter(http.StatusOK, "success", "user registered", userData)
	return c.JSON(http.StatusOK, response)
}
