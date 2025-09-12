// infra/persistence/repository/postgres_user_repository.go
package repository

import (
    "context"
    "user-crud/domain/model"
    "user-crud/infra/db"
)

type PostgresUserRepository struct {
    database *gorm.DB
}

func NewPostgresUserRepository() *PostgresUserRepository {
    return &PostgresUserRepository{database: db.DB}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
    err := r.database.WithContext(ctx).Create(&user).Error
    return user, err
}