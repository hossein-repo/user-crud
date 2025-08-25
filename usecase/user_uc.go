package usecase

import (
    "errors"
    "regexp"
    "user-crud/model"
    "user-crud/repository"

    "golang.org/x/crypto/bcrypt"
)

// ----------- Helpers -----------

func validateEmail(email string) bool {
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    return re.MatchString(email)
}

func validatePassword(password string) bool {
    return len(password) >= 6 // فقط طول ۶ کاراکتر برای شروع
}

// ----------- UseCases -----------

func RegisterUser(username, email, password string) error {
    if username == "" {
        return errors.New("username is required")
    }
    if !validateEmail(email) {
        return errors.New("invalid email format")
    }
    if !validatePassword(password) {
        return errors.New("password must be at least 6 characters")
    }

    // هش کردن پسورد
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return errors.New("could not hash password")
    }

    user := model.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
    }
    return repository.CreateUser(user)
}

func LoginUser(username, password string) (model.User, error) {
    u, err := repository.GetUser(username)
    if err != nil {
        return u, err
    }

    // بررسی صحت پسورد
    err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    if err != nil {
        return model.User{}, errors.New("invalid password")
    }
    return u, nil
}
