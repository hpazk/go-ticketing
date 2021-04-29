package report

type Services interface {
	FetchReport(creatorID, eventID uint, statusPayment string) (EventReport, error)
}

type services struct {
	repo repository
}

func ReportService() *services {
	repo := ReportRepository()
	return &services{repo}
}

func (s *services) FetchReport(creatorID, eventID uint, statusPayment string) (EventReport, error) {
	result, _ := s.repo.FetchReport(creatorID, eventID, statusPayment)
	report := EventReport{}

	report.WebinarDetil.TitleEvent = result[0].TitleEvent
	report.WebinarDetil.Description = result[0].Description
	report.WebinarDetil.LinkWebinar = result[0].LinkWebinar

	participants := func(report []Report) []EventParticipants {
		var participants []EventParticipants
		for _, r := range report {
			participant := EventParticipants{
				ParticipantEmail:    r.Email,
				ParticipantFullname: r.Fullname,
			}
			participants = append(participants, participant)
		}
		return participants
	}(result)
	report.Participants = participants

	totalAmount := func(report []Report) float64 {
		var total float64
		for _, r := range report {
			total += r.Amount
		}
		return total
	}(result)

	report.TotalAmount = totalAmount
	return report, nil
}
