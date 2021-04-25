package model

import (
	"gorm.io/gorm"
)

// User Entity
type User struct {
	gorm.Model
	Username    string
	Fullname    string
	Email       string
	Password    string
	Role        string        `sql:"type:ENUM('admin', 'participant', 'creator'); default:'participant'"`
	Event       []Event       `gorm:"foreignKey:CreatorID"`
	Transaction []Transaction `gorm:"foreignKey:ParticipantID"`
}
