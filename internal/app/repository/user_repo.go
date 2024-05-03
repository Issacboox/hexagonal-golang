package repository

import (
	m "bam/internal/app/model"

	"gorm.io/gorm"
)

type IRepository interface {
	CreateUser(user *m.User) error
	FindUserByID(id uint) (*m.User, error)
	UpdateUser(user *m.User) error
	DeleteUser(id uint) error
	FindUserByEmail(email string) (*m.User, error)
	FindUsers() ([]*m.User, error)
}

type UserRepository struct {
	db          *gorm.DB
	UserActions IRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *m.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindUserByID(id uint) (*m.User, error) {
	var user m.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *UserRepository) UpdateUser(user *m.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&m.User{}, id).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*m.User, error) {
	var user m.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindUsers() ([]*m.User, error) {
	var users []*m.User
	result := r.db.Find(&users)
	return users, result.Error
}
