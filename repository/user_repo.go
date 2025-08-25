package repository

import (
    "errors"
    "user-crud/model"
)

var users = map[string]model.User{}

func CreateUser(u model.User) error {
    if _, exists := users[u.Username]; exists {
        return errors.New("username already exists")
    }
    users[u.Username] = u
    return nil
}

func GetUser(username string) (model.User, error) {
    u, exists := users[username]
    if !exists {
        return model.User{}, errors.New("user not found")
    }
    return u, nil
}
