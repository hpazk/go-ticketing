package transaction

type request struct {
	ParticipantID uint    `json:"participant_id"`
	CreatorID     int     `json:"creator_id"`
	EventID       int     `json:"event_id"`
	Amount        float64 `json:"amount"`
	StatusPayment string  `json:"status_payment"`
}
