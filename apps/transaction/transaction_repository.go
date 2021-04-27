package transaction

import (
	"fmt"

	"github.com/hpazk/go-ticketing/apps/report"
	"github.com/hpazk/go-ticketing/database"
	"github.com/hpazk/go-ticketing/database/model"
	"gorm.io/gorm"
)

type repository interface {
	Store(tsx model.Transaction) (model.Transaction, error)
	Fetch() ([]model.Transaction, error)
	FindById(id uint) (model.Transaction, error)
	Update(tsx model.Transaction) (model.Transaction, error)
	Delete(id uint) error
	FindByEventID(eventID uint) ([]model.Transaction, error)
	FindByParticipant(participanID uint) (model.Transaction, error)
	FindDetil(transactionID uint) (report.Report, error)
}

type repo struct {
	db *gorm.DB
}

func TransactionRepository() *repo {
	db := database.GetDbInstance()
	return &repo{db}
}

func (r *repo) Store(tsx model.Transaction) (model.Transaction, error) {
	err := r.db.Create(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
}

func (r *repo) Fetch() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repo) FindById(id uint) (model.Transaction, error) {
	var tsx model.Transaction
	err := r.db.Where("id = ?", id).First(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
}

func (r *repo) FindByParticipant(participanID uint) (model.Transaction, error) {
	var tsx model.Transaction
	err := r.db.Where("participant_id = ?", participanID).First(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
}

func (r *repo) Update(transaction model.Transaction) (model.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repo) Delete(id uint) error {
	var tsx model.Transaction
	err := r.db.Where("id = ?", id).Delete(&tsx).Error
	if err != nil {
		return err
	}
	return nil
}

// Event - model.Transaction
func (r *repo) FindByEventID(eventID uint) ([]model.Transaction, error) {
	var transactions []model.Transaction

	// TODO WHERE transactions.status_payment = %s
	query := fmt.Sprintf("SELECT * FROM transactions JOIN users ON transactions.participant_id = users.id WHERE transactions.event_id = %d AND transactions.status_payment='passed';", eventID)
	err := r.db.Raw(query).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
func (r *repo) FindDetil(transactionID uint) (report.Report, error) {
	var transactions report.Report

	q := fmt.Sprintf(`SELECT users.fullname,
    users.email,
    events.id,
    events.title_event,
    events.description,
    events.link_webinar
	FROM users
	JOIN transactions ON users.id = transactions.participant_id
	JOIN events ON transactions.event_id = events.id
	WHERE transactions.id = %d;`, transactionID)

	err := r.db.Raw(q).Scan(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
