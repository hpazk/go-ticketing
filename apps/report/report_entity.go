package report

type ReportResult struct {
	Email         string
	EventID       int
	TransactionID int
	TitleEvent    string
	Amount        float64
}

type EventReport struct {
	TitleEvent       string
	TotalParticipant int
	Participants     []Participants
	TotalAmount      float64
}

type Participants struct {
	Email string
}
