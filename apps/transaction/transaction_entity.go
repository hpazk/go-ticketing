package transaction

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ParticipantID uint
	CreatorID     uint
	EventID       int
	StatusPayment string `gorm:"type:enum('passed', 'failed');''"`
	Amount        float64
}
