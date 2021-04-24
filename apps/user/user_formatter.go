package user

import (
	"time"

	"github.com/hpazk/go-booklib/database/model"
	"gorm.io/gorm"
)

// Request
type request struct {
	Username string `json:"username" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Response
type response struct {
	ID        uint           `json:"id"`
	Username  string         `json:"username"`
	Fullname  string         `json:"fullname"`
	Email     string         `json:"email"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	AuthToken string         `json:"auth_token"`
}

type loginResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	AuthToken string `json:"auth_token"`
}

func userLoginResponseFormatter(user model.User, authToken string) loginResponse {
	formatter := loginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
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
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		AuthToken: authToken,
	}

	return formatter
}
