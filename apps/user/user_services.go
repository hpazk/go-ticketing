package user

import (
	"errors"

	"github.com/hpazk/go-ticketing/database/model"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	signUp(req *request) (model.User, error)
	signIn(req *loginRequest) (model.User, error)
	FetchUsers() ([]model.User, error)
	FetchUserById(id uint) (model.User, error)
	FetchUserByRole(role string) (model.User, error)
	EditUser(id uint, req *updateRequest) error
	RemoveUser(id uint) error
	CheckExistEmail(email string) bool
	NewCreator(req *request) (model.User, error)
}

type services struct {
	repo repository
}

func UserService() *services {
	repo := UserRepository()
	return &services{repo}
}

func (s *services) signUp(req *request) (model.User, error) {
	user := model.User{}
	user.Username = req.Username
	user.Fullname = req.Fullname
	user.Email = req.Email
	user.Password = req.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashedPassword)

	newUser, err := s.repo.Store(user)
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

func (s *services) FetchUsers() ([]model.User, error) {
	var users []model.User
	users, err := s.repo.Fetch()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *services) FetchUserById(id uint) (model.User, error) {
	var user model.User
	user, err := s.repo.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *services) FetchUserByRole(role string) (model.User, error) {
	user, err := s.repo.FindByRole(role)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *services) EditUser(id uint, req *updateRequest) error {
	user, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	user.Username = req.Username
	user.Fullname = req.Fullname
	user.Email = req.Email

	err = s.repo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *services) RemoveUser(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

// Creator
func (s *services) NewCreator(req *request) (model.User, error) {
	user := model.User{}
	user.Username = req.Username
	user.Fullname = req.Fullname
	user.Email = req.Email
	user.Password = req.Password
	user.Role = "creator"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashedPassword)

	newUser, err := s.repo.Store(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
