package service

import (
	"errors"
	m "bam/internal/app/model"
	r "bam/internal/app/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo r.IRepository
}

func NewUserService(repo r.IRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id uint) (*m.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *UserService) CreateUser(user *m.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)

}

func (s *UserService) UpdateUser(user *m.User) error {
	return s.repo.UpdateUser(user)

}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) AuthenticateUser(email, password string) (*m.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetUsers() ([]*m.User, error) {
	return s.repo.FindUsers()
}
