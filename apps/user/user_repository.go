package user

import (
	"github.com/jinzhu/gorm"
)

type userRepository interface {
	Store(user User) (User, error)
	Fetch() ([]User, error)
	Update(user User) (User, error)
	FindById(id uint) (User, error)
	FindByEmail(email string) (User, error)
	// Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save New User
func (r *repository) Store(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get All Users
func (r *repository) Fetch() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

// Get User by Id
func (r *repository) FindById(id uint) (User, error) {
	var user User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get User By Email
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update user
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// // Delete User
// func (r *repository) Delete(id uint) error {
// 	usersStorage[id-1] = usersStorage[len(usersStorage)-1]
// 	usersStorage[uint(len(usersStorage))-1] = User{}
// 	usersStorage = usersStorage[:uint(len(usersStorage))-1]
// 	return nil
// }
