package transaction

type response struct {
	ID            uint    `json:"id"`
	ParticipantID uint    `json:"participan_id"`
	CreatorID     uint    `json:"creator_id"`
	EventID       int     `json:"event_id"`
	StatusPayment string  `json:"status_payment"`
	Amount        float64 `json:"amount"`
}

func tsxFormatter(tsx Transaction) response {
	formatter := response{
		ID:            tsx.ID,
		ParticipantID: tsx.ParticipantID,
		CreatorID:     tsx.CreatorID,
		EventID:       tsx.EventID,
		StatusPayment: tsx.StatusPayment,
		Amount:        tsx.Amount,
	}

	return formatter
}

func tsxsFormatter(tsxs []Transaction) []response {
	formatter := []response{}

	for _, campaign := range tsxs {
		c := tsxFormatter(campaign)
		formatter = append(formatter, c)
	}
	return formatter
}
