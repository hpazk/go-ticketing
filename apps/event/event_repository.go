package event

import (
	"fmt"

	"github.com/hpazk/go-booklib/database/model"
	"gorm.io/gorm"
)

type repository interface {
	Store(event model.Event) (model.Event, error)
	Fetch() ([]model.Event, error)
	FindById(id uint) (model.Event, error)
	Update(event model.Event) (model.Event, error)
	Delete(id uint) error
	FetchReport(creatorID uint) ([]model.User, error)
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
func (r *repo) FetchReport(creatorID uint) ([]model.User, error) {
	var report []model.User
	q := fmt.Sprintf("SELECT * FROM transactions JOIN users ON transactions.participant_id = users.id JOIN events ON transactions.event_id = events.id WHERE events.creator_id = %d;", creatorID)
	err := r.db.Raw(q).Scan(&report).Error
	if err != nil {
		return report, err
	}

	fmt.Println(report)

	return report, nil
}
