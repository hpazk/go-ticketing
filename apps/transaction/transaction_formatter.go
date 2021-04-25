package transaction

import "github.com/hpazk/go-ticketing/database/model"

type request struct {
	EventID uint    `json:"event_id" validate:"required"`
	Amount  float64 `json:"amount"`
}

type response struct {
	ID            uint    `json:"id"`
	ParticipantID uint    `json:"participan_id"`
	CreatorID     uint    `json:"creator_id"`
	EventID       uint    `json:"event_id"`
	StatusPayment string  `json:"status_payment"`
	Amount        float64 `json:"amount"`
}

func tsxFormatter(tsx model.Transaction) response {
	formatter := response{
		ID:            tsx.ID,
		ParticipantID: tsx.ParticipantID,
		EventID:       tsx.EventID,
		StatusPayment: tsx.StatusPayment,
		Amount:        tsx.Amount,
	}

	return formatter
}

func tsxsFormatter(tsxs []model.Transaction) []response {
	formatter := []response{}

	for _, campaign := range tsxs {
		c := tsxFormatter(campaign)
		formatter = append(formatter, c)
	}
	return formatter
}
