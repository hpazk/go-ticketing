package transaction

import (
	"github.com/hpazk/go-booklib/apps/event"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ParticipantID uint
	CreatorID     int
	EventID       int
	Amount        float64
	Event         []event.Event
}
