package report

type Services interface {
	FetchEventParticipant(creatorID uint, statusPayment string) ([]Report, error)
}

type services struct {
	repo repository
}

func ReportService() *services {
	repo := ReportRepository()
	return &services{repo}
}

func (s *services) FetchEventParticipant(creatorID uint, statusPayment string) ([]Report, error) {
	// eventReport := EventReport{}

	report, err := s.repo.FetchTransactionReport(creatorID, statusPayment)
	if err != nil {
		return report, nil
	}
	// if err != nil {
	// 	return eventReport, nil
	// }

	// participants := func(report []Report) []EventParticipants {
	// 	var participants []EventParticipants
	// 	for _, r := range report {
	// 		participant := EventParticipants{
	// 			ParticipantEmail:    r.Email,
	// 			ParticipantFullname: r.Fullname,
	// 			StatusPayment:       r.StatusPayment,
	// 		}
	// 		participants = append(participants, participant)
	// 	}
	// 	return participants
	// }(report)

	// eventReport.Participants = participants
	// return eventReport, nil
	return report, nil
}
