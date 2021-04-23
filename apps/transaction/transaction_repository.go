package transaction

import (
	"fmt"

	"github.com/hpazk/go-booklib/database/model"
	"gorm.io/gorm"
)

type repository interface {
	Store(tsx model.Transaction) (model.Transaction, error)
	Fetch() ([]model.Transaction, error)
	FindById(id uint) (model.Transaction, error)
	Update(tsx model.Transaction) (model.Transaction, error)
	Delete(id uint) error
	FindByEventID(eventID uint) ([]model.Transaction, error)
}

type repo struct {
	db *gorm.DB
}

func transactionRepository(db *gorm.DB) *repo {
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
	var tsxs []model.Transaction
	err := r.db.Find(&tsxs).Error
	if err != nil {
		return tsxs, err
	}

	fmt.Println(tsxs)

	return tsxs, nil
}

func (r *repo) FindById(id uint) (model.Transaction, error) {
	var tsx model.Transaction
	err := r.db.Find(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
}

func (r *repo) Update(tsx model.Transaction) (model.Transaction, error) {
	err := r.db.Save(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
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
