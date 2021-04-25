package event

import "time"

type request struct {
	CreatorID         uint      `json:"creator_id" validate:"required"`
	TitleEvent        string    `json:"title_event" validate:"required"`
	LinkWebinar       string    `json:"link_webinar"`
	Description       string    `json:"description" validate:"required"`
	Banner            string    `json:"banner" validate:"required"`
	Price             float64   `json:"price" validate:"required"`
	Quantity          int       `json:"quantity" validate:"required"`
	Status            string    `json:"status" validate:"required"`
	EventStartDate    time.Time `json:"event_start_date" validate:"required"`
	EventEndDate      time.Time `json:"event_end_date" validate:"required"`
	CampaignStartDate time.Time `json:"capaign_start_date" validate:"required"`
	CampaignEndDate   time.Time `json:"campaign_end_date" validate:"required"`
}
