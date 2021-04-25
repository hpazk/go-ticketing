package user

import (
	"errors"

	"github.com/hpazk/go-ticketing/database/model"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	signUp(req *request) (model.User, error)
	signIn(req *loginRequest) (model.User, error)
	// FetchUsers() ([]model.User, error)
	FetchUserById(id uint) (model.User, error)
	// FetchUserByEmail(email string) (model.User, error)
	// UpdateUser(id uint, req *updateRequest) (model.User, error)
	// DeleteUser(id uint) error
	CheckExistEmail(email string) bool
}

type services struct {
	repo repository
}

func UserService() *services {
	repo := UserRepository()
	return &services{repo}
}

func (s *services) signUp(req *request) (model.User, error) {
	userReg := model.User{}
	userReg.Username = req.Username
	userReg.Fullname = req.Fullname
	userReg.Email = req.Email
	userReg.Password = req.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return userReg, err
	}

	userReg.Password = string(hashedPassword)

	newUser, err := s.repo.Store(userReg)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *services) signIn(req *loginRequest) (model.User, error) {
	email := req.Email
	password := req.Password

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil

}
func (s *services) CheckExistEmail(email string) bool {
	if _, err := s.repo.FindByEmail(email); err != nil {
		return false
	}

	return true
}

// // TODO error-handling
// func (s *services) FetchUsers() ([]model.User, error) {
// 	var users []model.User
// 	users, err := s.repo.Fetch()
// 	if err != nil {
// 		return users, err
// 	}

// 	return users, nil
// }

func (s *services) FetchUserById(id uint) (model.User, error) {
	var user model.User
	user, err := s.repo.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (s *services) FetchUserByEmail(email string) (model.User, error) {
// 	var user model.User
// 	user, err := s.repo.FindByEmail(email)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (s *services) UpdateUser(id uint, req *updateRequest) (model.User, error) {
// 	userReg := model.User{}
// 	userReg.ID = id
// 	userReg.Name = req.Name
// 	userReg.Address = req.Address
// 	// userReg.Photo = ""
// 	userReg.Email = req.Email
// 	// userReg.Role = ""
// 	userReg.CreatedAt = time.Now()
// 	userReg.UpdatedAt = time.Now()

// 	editedUser, err := s.repo.Update(userReg)
// 	if err != nil {
// 		return editedUser, err
// 	}

// 	return editedUser, nil
// }

// func (s *services) DeleteUser(id uint) error {
// 	if err := s.repo.Delete(id); err != nil {
// 		return err
// 	}

// 	return nil
// }
