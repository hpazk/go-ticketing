package event

import (
	"time"

	"github.com/hpazk/go-ticketing/database/model"
)

type request struct {
	CreatorID         uint      `json:"creator_id" validate:"required"`
	TitleEvent        string    `json:"title_event" validate:"required"`
	LinkWebinar       string    `json:"link_webinar"`
	Description       string    `json:"description" validate:"required"`
	Banner            string    `json:"banner"`
	Price             float64   `json:"price"`
	Quantity          int       `json:"quantity"`
	Status            string    `json:"status"`
	EventStartDate    time.Time `json:"event_start_date"`
	EventEndDate      time.Time `json:"event_end_date"`
	CampaignStartDate time.Time `json:"capaign_start_date"`
	CampaignEndDate   time.Time `json:"campaign_end_date"`
}

type response struct {
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

type updateRequest struct {
	CreatorID         uint      `json:"creator_id" validate:"required"`
	TitleEvent        string    `json:"title_event" validate:"required"`
	LinkWebinar       string    `json:"link_webinar" validate:"required"`
	Description       string    `json:"description" validate:"required"`
	Banner            string    `json:"banner" validate:"required"`
	Price             float64   `json:"price" validate:"required"`
	Quantity          int       `json:"quantity" validate:"required"`
	Status            string    `json:"status" validate:"required"`
	EventStartDate    time.Time `json:"event_start_date"`
	EventEndDate      time.Time `json:"event_end_date"`
	CampaignStartDate time.Time `json:"capaign_start_date"`
	CampaignEndDate   time.Time `json:"campaign_end_date"`
}

func eventResponseFormatter(event model.Event) response {
	formatter := response{
		ID:                event.ID,
		CreatorID:         event.CreatorID,
		TitleEvent:        event.TitleEvent,
		LinkWebinar:       event.LinkWebinar,
		Description:       event.Description,
		Banner:            event.Banner,
		Price:             event.Price,
		Quantity:          event.Quantity,
		Status:            event.Status,
		EventStartDate:    event.EventStartDate,
		EventEndDate:      event.EventEndDate,
		CampaignStartDate: event.CampaignStartDate,
		CampaignEndDate:   event.CampaignStartDate,
	}

	return formatter
}

func eventsResponseFormatter(events []model.Event) []response {
	formatter := []response{}

	for _, event := range events {
		c := eventResponseFormatter(event)
		formatter = append(formatter, c)
	}
	return formatter
}
