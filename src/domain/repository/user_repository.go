// src/domain/repository/user_repository.go
package repository

import (
    "context"
    "user-crud/domain/model"
)

type UserRepository interface {
    Create(ctx context.Context, user *model.User) error
    GetByID(ctx context.Context, id uint) (*model.User, error)
    GetByUsername(ctx context.Context, username string) (*model.User, error)
    GetByEmail(ctx context.Context, email string) (*model.User, error)
    GetByMobile(ctx context.Context, mobileNumber string) (*model.User, error)
    Update(ctx context.Context, user *model.User) error
    Delete(ctx context.Context, id uint) error
    ExistsUsername(ctx context.Context, username string) (bool, error)
    ExistsEmail(ctx context.Context, email string) (bool, error)
    ExistsMobile(ctx context.Context, mobileNumber string) (bool, error)
}