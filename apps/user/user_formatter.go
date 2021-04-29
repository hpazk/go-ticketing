package user

import (
	"github.com/hpazk/go-ticketing/database/model"
)

// Request
type request struct {
	Username string `json:"username" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type updateRequest struct {
	Username string `json:"username" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// Response
type response struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	AuthToken string `json:"auth_token"`
}

type loginResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	AuthToken string `json:"auth_token"`
}

type basicRsponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func userLoginResponseFormatter(user model.User, authToken string) loginResponse {
	formatter := loginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Role:      user.Role,
		AuthToken: authToken,
	}

	return formatter
}

func userResponseFormatter(user model.User, authToken string) response {
	formatter := response{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Role:      user.Role,
		AuthToken: authToken,
	}

	return formatter
}

func userBasicResponseFormatter(user model.User) basicRsponse {
	formatter := basicRsponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
		Role:     user.Role,
	}

	return formatter
}

func usersBasicResponseFormatter(users []model.User) []basicRsponse {
	formatter := []basicRsponse{}

	for _, user := range users {
		c := userBasicResponseFormatter(user)
		formatter = append(formatter, c)
	}
	return formatter
}
