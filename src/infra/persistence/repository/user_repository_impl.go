// src/infra/persistence/repository/user_repository_impl.go
package repository

import (
	"context"
	"user-crud/domain/model"
	"user-crud/infra/db"

	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{database: db.DB}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.database.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.database.WithContext(ctx).
		Where("username = ?", username).
		First(&user).Error
	return &user, err
}

// بقیه متدها...
