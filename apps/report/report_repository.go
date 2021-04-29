package report

import (
	"fmt"

	"github.com/hpazk/go-ticketing/database"
	"gorm.io/gorm"
)

type repository interface {
	FetchReport(creatorID, eventID uint, statusPayment string) ([]Report, error)
}

type repo struct {
	db *gorm.DB
}

func ReportRepository() *repo {
	db := database.GetDbInstance()
	return &repo{db}
}

func (r *repo) FetchReport(creatorID, eventID uint, statusPayment string) ([]Report, error) {
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
    AND transactions.status_payment = '%s';`, creatorID, eventID, statusPayment)

	err := r.db.Raw(q).Scan(&report).Error
	if err != nil {
		return report, err
	}

	return report, nil
}
