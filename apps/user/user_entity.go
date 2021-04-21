package user

import (
	"github.com/hpazk/go-booklib/apps/event"
	"github.com/hpazk/go-booklib/apps/transaction"
	"gorm.io/gorm"
)

// User Entity
type User struct {
	gorm.Model
	Username    string
	Fullname    string
	Email       string
	Password    string
	Role        string
	Event       []event.Event             `gorm:"foreignKey:CreatorID"`
	Transaction []transaction.Transaction `gorm:"foreignKey:ParticipantID"`
}
