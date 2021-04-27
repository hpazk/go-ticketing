package report

type Services interface {
	// FetchEventReport(creatorID uint, statusPayment string) (EventReport, error)
	FetchReport(creatorID, eventID uint) (EventReport, error)
}

type services struct {
	repo repository
}

func ReportService() *services {
	repo := ReportRepository()
	return &services{repo}
}

func (s *services) FetchReport(creatorID, eventID uint) (EventReport, error) {
	result, _ := s.repo.FetchReport(creatorID, eventID)
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
	return report, nil
}

// func (s *services) FetchEventReport(creatorID uint, statusPayment string) (EventReport, error) {
// 	eventReport := EventReport{}

// 	report, err := s.repo.FetchReport(creatorID, statusPayment)
// 	if err != nil {
// 		return eventReport, nil
// 	}

// 	// eventReport.TitleEvent = report[0].TitleEvent
// 	eventReport.WebinarDetil.TitleEvent = report[0].TitleEvent
// 	eventReport.WebinarDetil.Description = report[0].Description
// 	eventReport.WebinarDetil.LinkWebinar = report[0].LinkWebinar

// 	participants := func(report []Report) []EventParticipants {
// 		var participants []EventParticipants
// 		for _, r := range report {
// 			participant := EventParticipants{
// 				ParticipantEmail:    r.Email,
// 				ParticipantFullname: r.Fullname,
// 				StatusPayment:       r.StatusPayment,
// 			}
// 			participants = append(participants, participant)
// 		}
// 		return participants
// 	}(report)

// 	eventReport.Participants = participants
// 	eventReport.TotalAmount = func(report []Report) float64 {
// 		var totalAmount float64
// 		for _, r := range report {
// 			totalAmount += r.Amount
// 		}
// 		return totalAmount
// 	}(report)
// 	return eventReport, nil
// 	// return report, nil
// }
