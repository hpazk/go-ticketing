package event

import (
	"github.com/hpazk/go-ticketing/database"
	"github.com/hpazk/go-ticketing/database/model"
	"gorm.io/gorm"
)

type repository interface {
	Store(event model.Event) (model.Event, error)
	Fetch() ([]model.Event, error)
	FindById(id uint) (model.Event, error)
	Update(event model.Event) (model.Event, error)
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func EventRepository() *repo {
	db := database.GetDbInstance()
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
