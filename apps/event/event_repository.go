package event

import (
	"fmt"

	"github.com/hpazk/go-booklib/apps/report"
	"github.com/hpazk/go-booklib/database/model"
	"gorm.io/gorm"
)

type repository interface {
	Store(event model.Event) (model.Event, error)
	Fetch() ([]model.Event, error)
	FindById(id uint) (model.Event, error)
	Update(event model.Event) (model.Event, error)
	Delete(id uint) error
	FetchReport(creatorID uint, eventID uint) ([]report.ReportResult, error)
}

type repo struct {
	db *gorm.DB
}

func eventRepository(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) Store(event model.Event) (model.Event, error) {
	err := r.db.Create(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repo) Fetch() ([]model.Event, error) {
	var events []model.Event
	err := r.db.Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repo) FindById(id uint) (model.Event, error) {
	var event model.Event
	err := r.db.Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repo) Update(event model.Event) (model.Event, error) {
	err := r.db.Save(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repo) Delete(id uint) error {
	var event model.Event
	err := r.db.Where("id = ?", id).Delete(&event).Error
	if err != nil {
		return err
	}
	return nil
}

// model.Event - Transaction - User
func (r *repo) FetchReport(creatorID uint, eventID uint) ([]report.ReportResult, error) {
	var result []report.ReportResult
	// TODO events.id = eventID
	q := fmt.Sprintf("SELECT transactions.id AS transaction_id,events.id AS event_id,users.id AS user_id,users.email,events.title_event,events.price,transactions.status_payment,transactions.amount FROM events JOIN transactions ON events.id = transactions.event_id JOIN users ON transactions.participant_id = users.id WHERE events.creator_id = %d AND events.id = %d", creatorID, eventID)
	err := r.db.Raw(q).Scan(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
