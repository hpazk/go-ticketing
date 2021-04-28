package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/hpazk/go-ticketing/auth"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userServices UserServices
	authServices auth.AuthServices
}

func UserHandler(userServices UserServices, authServices auth.AuthServices) *userHandler {
	return &userHandler{userServices, authServices}
}

func (h *userHandler) PostUserRegistration(c echo.Context) error {
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
	newUser, err := h.userServices.signUp(req)
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

func (h *userHandler) PostUserLogin(c echo.Context) error {
	req := new(loginRequest)

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

	signedInUser, err := h.userServices.signIn(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	authToken, err := h.authServices.GetAccessToken(signedInUser.ID, signedInUser.Role, signedInUser.Email)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", "something went wrong", nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	userData := userLoginResponseFormatter(signedInUser, authToken)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user authenticated", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) PostUserLogout(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.M{"message": "user-logout"})
}

func (h *userHandler) GetUsers(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	if findByRole := c.QueryParam("role"); findByRole != "" {
		role := c.QueryParam("role")

		user, err := h.userServices.FetchUserByRole(role)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.M{"message": err.Error()})
		} else {
			response := userBasicResponseFormatter(user)
			return c.JSON(http.StatusOK, response)
		}
	}

	users, err := h.userServices.FetchUsers()
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "fail", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	usersData := usersBasicResponseFormatter(users)

	response := helper.ResponseFormatter(http.StatusOK, "success", "all user successfully fetched", usersData)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUser(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	response, _ := h.userServices.FetchUserById(uint(id))
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) PatchUser(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
		response := helper.ResponseFormatter(http.StatusUnauthorized, "fail", "Please provide valid credentials", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request)

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

	editedUser, err := h.userServices.EditUser(uint(id), req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "fail", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userData := userBasicResponseFormatter(editedUser)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user successfully updated", userData)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	accessToken := c.Get("user").(*jwt.Token)
	claims := accessToken.Claims.(jwt.MapClaims)
	role := claims["user_role"]

	if role != "admin" {
		response := helper.ResponseFormatterWD(http.StatusUnauthorized, "fail", "Please provide valid credentials")
		return c.JSON(http.StatusUnauthorized, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.userServices.RemoveUser(uint(id)); err != nil {
		response := helper.ResponseFormatterWD(http.StatusInternalServerError, "fail", err)
		return c.JSON(http.StatusOK, response)
	}
	message := fmt.Sprintf("user %d was deleted", id)
	response := helper.ResponseFormatterWD(http.StatusOK, "success", message)
	return c.JSON(http.StatusOK, response)
}
