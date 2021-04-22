package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	signUp(req *request) (User, error)
	signIn(req *loginRequest) (User, error)
	// FetchUsers() ([]User, error)
	FetchUserById(id uint) (User, error)
	// FetchUserByEmail(email string) (User, error)
	// UpdateUser(id uint, req *updateRequest) (User, error)
	// DeleteUser(id uint) error
	CheckExistEmail(email string) bool
}

type services struct {
	repo repository
}

func UserService(repo repository) *services {
	return &services{repo}
}

func (s *services) signUp(req *request) (User, error) {
	userReg := User{}
	userReg.Username = req.Username
	userReg.Fullname = req.Fullname
	userReg.Email = req.Email
	userReg.Password = req.Password
	userReg.Role = "member"
	userReg.CreatedAt = time.Now()
	userReg.UpdatedAt = time.Now()

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

// TODO user-login
func (s *services) signIn(req *loginRequest) (User, error) {
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
// func (s *services) FetchUsers() ([]User, error) {
// 	var users []User
// 	users, err := s.repo.Fetch()
// 	if err != nil {
// 		return users, err
// 	}

// 	return users, nil
// }

func (s *services) FetchUserById(id uint) (User, error) {
	var user User
	user, err := s.repo.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (s *services) FetchUserByEmail(email string) (User, error) {
// 	var user User
// 	user, err := s.repo.FindByEmail(email)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (s *services) UpdateUser(id uint, req *updateRequest) (User, error) {
// 	userReg := User{}
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
