package usecase

import (
    "errors"
    "user-crud/model"
    "user-crud/repository"
)

func RegisterUser(username, email, password string) error {
    user := model.User{
        Username: username,
        Email:    email,
        Password: password,
    }
    return repository.CreateUser(user)
}

func LoginUser(username, password string) (model.User, error) {
    u, err := repository.GetUser(username)
    if err != nil {
        return u, err
    }
    if u.Password != password {
        return model.User{}, errors.New("invalid password")
    }
    return u, nil
}
