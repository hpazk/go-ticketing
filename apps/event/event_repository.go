package event

import "gorm.io/gorm"

type repository interface {
	Store(event Event) (Event, error)
	Fetch() ([]Event, error)
	FindById(id uint) (Event, error)
	Update(event Event) (Event, error)
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func eventRepository(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) Store(event Event) (Event, error) {
	err := r.db.Create(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repo) Fetch() ([]Event, error) {
	var events []Event
	err := r.db.Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repo) FindById(id uint) (Event, error) {
	var event Event
	err := r.db.Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repo) Update(event Event) (Event, error) {
	err := r.db.Save(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repo) Delete(id uint) error {
	var event Event
	err := r.db.Where("id = ?", id).Delete(&event).Error
	if err != nil {
		return err
	}
	return nil
}
