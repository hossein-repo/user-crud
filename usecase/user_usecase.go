package usecase

import (
	"errors"
	"time"
	"user-crud/infra/db"
	"user-crud/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, email, password string) error {
    var user model.User
    if err := db.DB.Where("username = ?", username).Or("email = ?", email).First(&user).Error; err == nil {
        return errors.New("username or email already exists")
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    newUser := model.User{
        Username: username,
        Email:    email,
        Password: string(hashed),
        Enabled:  true,
        CreatedBy: 0, // admin یا سیستم
    }

    return db.DB.Create(&newUser).Error
}

func LoginUser(username, password string) (string, error) {
    var user model.User
    if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
        return "", errors.New("user not found")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  user.ID,
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
