package report

type Report struct {
	TitleEvent    string
	Description   string
	LinkWebinar   string
	Email         string
	Fullname      string
	StatusPayment string
}
type EventParticipants struct {
	ParticipantEmail    string `json:"participan_email"`
	ParticipantFullname string `json:"participan_fullname"`
	StatusPayment       string `json:"status_payment"`
}

type EventReport struct {
	TitleEvent   string              `json:"title_event"`
	Participants []EventParticipants `json:"participants"`
}
