package event

import "time"

type request struct {
	ID                uint      `json:"id"`
	CreatorID         uint      `json:"creator_id"`
	TitleEvent        string    `json:"title_event"`
	LinkWebinar       string    `json:"link_webinar"`
	Description       string    `json:"description"`
	Banner            string    `json:"banner"`
	Price             float64   `json:"price"`
	Quantity          int       `json:"quantity"`
	Status            string    `json:"status"`
	EventStartDate    time.Time `json:"event_start_date"`
	EventEndDate      time.Time `json:"event_end_date"`
	CampaignStartDate time.Time `json:"capaign_start_date"`
	CampaignEndDate   time.Time `json:"campaign_end_date"`
}
