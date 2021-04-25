package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ParticipantID uint
	EventID       uint
	ImagePath     string
	StatusPayment string `sql:"type:ENUM('passed', 'failed')" gorm:"default:failed"`
	Amount        float64
}
