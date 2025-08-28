package repository

import (
	"errors"
	"user-crud/infra/db"
	"user-crud/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetByID(id uint) (*model.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *model.User) error {
	return db.DB.Create(user).Error
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := db.DB.Preload("UserRoles").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := db.DB.Preload("UserRoles").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
