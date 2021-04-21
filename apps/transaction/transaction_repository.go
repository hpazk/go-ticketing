package transaction

import "gorm.io/gorm"

type repository interface {
	Store(tsx Transaction) (Transaction, error)
	Fetch() ([]Transaction, error)
	FindById(id uint) (Transaction, error)
	Update(tsx Transaction) (Transaction, error)
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func transactionRepository(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) Store(tsx Transaction) (Transaction, error) {
	err := r.db.Create(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
}

func (r *repo) Fetch() ([]Transaction, error) {
	var events []Transaction
	err := r.db.Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repo) FindById(id uint) (Transaction, error) {
	var tsx Transaction
	err := r.db.Find(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
}

func (r *repo) Update(tsx Transaction) (Transaction, error) {
	err := r.db.Save(&tsx).Error
	if err != nil {
		return tsx, err
	}

	return tsx, nil
}

func (r *repo) Delete(id uint) error {
	var tsx Transaction
	err := r.db.Where("id = ?", id).Delete(&tsx).Error
	if err != nil {
		return err
	}
	return nil
}
