package report

type Report struct {
	TitleEvent    string
	Description   string
	LinkWebinar   string
	Email         string
	Fullname      string
	StatusPayment string
	Amount        float64
}
type EventParticipants struct {
	ParticipantEmail    string `json:"participan_email"`
	ParticipantFullname string `json:"participan_fullname"`
	StatusPayment       string `json:"status_payment"`
}

type EventReport struct {
	WebinarDetil WebinarDetil        `json:"webinar_detil"`
	TotalAmount  float64             `json:"total_amount"`
	Participants []EventParticipants `json:"participants"`
}

type WebinarDetil struct {
	TitleEvent  string
	Description string
	LinkWebinar string
}
