package event

import (
	"time"

	"github.com/hpazk/go-booklib/apps/transaction"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	CreatorID         uint
	TitleEvent        string
	LinkWebinar       string
	Description       string
	Banner            string
	Price             float64
	Quantity          int
	Status            string `gorm:"type:enum('draft', 'release'); default:'draft'"`
	EventStartDate    time.Time
	EventEndDate      time.Time
	CampaignStartDate time.Time
	CampaignEndDate   time.Time
	Transaction       []transaction.Transaction
}
