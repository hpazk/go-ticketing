package transaction

import (
	"github.com/hpazk/go-booklib/apps/event"
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	ParticipantID uint
	CreatorID     int
	EventID       int
	Amount        float64
	Event         []event.Event
}
