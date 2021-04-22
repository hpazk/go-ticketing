package model

import (
	"gorm.io/gorm"
)

// User Entity
type User struct {
	gorm.Model
	Username string
	Fullname string
	Email    string
	Password string
	Role     string
}
