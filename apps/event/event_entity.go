package event

import (
	"time"

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
	Status            string
	EventStartDate    time.Time
	EventEndDate      time.Time
	CampaignStartDate time.Time
	CampaignEndDate   time.Time
}
