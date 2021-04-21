package user

import (
	"gorm.io/gorm"
)

type repository interface {
	Store(user User) (User, error)
	Fetch() ([]User, error)
	Update(user User) (User, error)
	FindById(id uint) (User, error)
	FindByEmail(email string) (User, error)
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func userRepository(db *gorm.DB) *repo {
	return &repo{db}
}

// Save New User
func (r *repo) Store(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get All Users
func (r *repo) Fetch() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

// Get User by Id
func (r *repo) FindById(id uint) (User, error) {
	var user User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get User By Email
func (r *repo) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update user
func (r *repo) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Delete User
func (r *repo) Delete(id uint) error {
	var user User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
