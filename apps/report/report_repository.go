package report

import (
	"fmt"

	"github.com/hpazk/go-ticketing/database"
	"github.com/hpazk/go-ticketing/database/model"
	"gorm.io/gorm"
)

type repository interface {
	FetchReport(creatorID, eventID uint) ([]Report, error)
	// Fetch(creatorID uint) ([]model.User, error)
}

type repo struct {
	db *gorm.DB
}

func ReportRepository() *repo {
	db := database.GetDbInstance()
	return &repo{db}
}

func (r *repo) FetchReport(creatorID, eventID uint) ([]Report, error) {
	var report []Report

	q := fmt.Sprintf(`SELECT users.fullname,
    users.email,
    events.title_event,
	events.description,
	events.link_webinar,
	events.price,
	transactions.status_payment,
	transactions.amount
	FROM transactions
    JOIN events ON transactions.event_id = events.id
    JOIN users ON transactions.participant_id = users.id
	WHERE events.creator_id = %d
	AND events.id = %d
    AND transactions.status_payment = 'passed';`, creatorID, eventID)

	err := r.db.Raw(q).Scan(&report).Error
	if err != nil {
		return report, err
	}

	fmt.Println(report)
	return report, nil
}

func (r *repo) Fetch() ([]model.User, error) {
	users := []model.User{}
	err := r.db.Preload("Transaction").Preload("Event").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
