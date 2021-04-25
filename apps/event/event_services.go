package event

import (
	"github.com/hpazk/go-booklib/apps/report"
	"github.com/hpazk/go-booklib/database/model"
)

type Services interface {
	SaveEvent(req *request) (model.Event, error)
	FetchEvents() ([]model.Event, error)
	FetchEvent(id uint) (model.Event, error)
	EditEvent(id uint, req *request) (model.Event, error)
	RemoveEvent(id uint) error
	FetchEventReport(creatorID uint, eventID uint) (report.EventReport, error)
}

type services struct {
	repo repository
}

func EventService() *services {
	repo := EventRepository()
	return &services{repo}
}

func (s *services) SaveEvent(req *request) (model.Event, error) {
	var event model.Event
	event.CreatorID = req.CreatorID
	event.TitleEvent = req.TitleEvent
	event.LinkWebinar = req.LinkWebinar
	event.Description = req.Description
	event.Banner = req.Banner
	event.Price = req.Price
	event.Quantity = req.Quantity
	event.Status = req.Status
	event.EventStartDate = req.EventStartDate
	event.EventEndDate = req.EventEndDate
	event.CampaignStartDate = req.CampaignStartDate
	event.CampaignEndDate = req.CampaignEndDate

	savedEvent, err := s.repo.Store(event)
	if err != nil {
		return savedEvent, nil
	}

	return savedEvent, nil
}

func (s *services) FetchEvents() ([]model.Event, error) {
	events, err := s.repo.Fetch()
	if err != nil {
		return events, nil
	}

	return events, nil
}

func (s *services) FetchEvent(id uint) (model.Event, error) {
	event, err := s.repo.FindById(id)
	if err != nil {
		return event, nil
	}
	return event, nil
}

func (s *services) EditEvent(id uint, req *request) (model.Event, error) {
	var event model.Event
	event.ID = id
	event.CreatorID = req.CreatorID
	event.TitleEvent = req.TitleEvent
	event.LinkWebinar = req.LinkWebinar
	event.Description = req.Description
	event.Banner = req.Banner
	event.Price = req.Price
	event.Quantity = req.Quantity
	event.Status = req.Status
	event.EventStartDate = req.EventStartDate
	event.EventEndDate = req.EventEndDate
	event.CampaignStartDate = req.CampaignStartDate
	event.CampaignEndDate = req.CampaignEndDate

	editedEvent, err := s.repo.Update(event)

	if err != nil {
		return editedEvent, nil
	}
	return editedEvent, nil
}

func (s *services) RemoveEvent(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *services) FetchEventReport(creatorID uint, eventID uint) (report.EventReport, error) {
	result, _ := s.repo.FetchReport(creatorID, eventID)
	eventReport := report.EventReport{}
	eventReport.TitleEvent = result[0].TitleEvent
	eventReport.TotalParticipant = len(result)
	eventReport.Participants = func(res []report.ReportResult) []report.Participants {
		participants := []report.Participants{}
		for _, r := range res {
			participant := report.Participants{}
			participant.Email = r.Email
			participants = append(participants, participant)
		}
		return participants
	}(result)
	eventReport.TotalAmount = func(res []report.ReportResult) float64 {
		var totalAmount float64
		for _, r := range res {
			totalAmount += r.Amount
		}
		return totalAmount
	}(result)
	return eventReport, nil
}
