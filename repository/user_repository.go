package repository

import (
    "user-crud/infra/db"
    "user-crud/model"
)

type UserRepository interface {
    Create(user *model.User) error
    GetByUsername(username string) (*model.User, error)
    GetByID(id uint) (*model.User, error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
    return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
    return db.DB.Create(user).Error
}

func (r *userRepo) GetByUsername(username string) (*model.User, error) {
    var user model.User
    if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) GetByID(id uint) (*model.User, error) {
    var user model.User
    if err := db.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
