// src/usecase/user_usecase.go
package usecase

import (
	"context"
	"user-crud/domain/model"
	"user-crud/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, username, password string) (*model.User, error)
	GetProfile(ctx context.Context, userID uint) (*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uc *userUsecase) Register(ctx context.Context, user *model.User) error {
	// هش کردن پسورد
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return uc.userRepo.Create(ctx, user)
}

func (uc *userUsecase) Login(ctx context.Context, username, password string) (*model.User, error) {
	user, err := uc.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// بررسی پسورد
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
