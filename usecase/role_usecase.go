package usecase

import (
    "errors"
    "user-crud/model"
    "user-crud/repository"
)

var roleRepo = repository.NewRoleRepository()
var userRoleRepo = repository.NewUserRoleRepository()

func AssignRoleToUser(userID, roleID uint) error {
    _, err := roleRepo.GetByID(roleID)
    if err != nil {
        return err
    }
    userRole := &model.UserRole{
        UserID: userID,
        RoleID: roleID,
    }
    return userRoleRepo.Create(userRole)
}

func GetUserRoles(userID uint) ([]model.Role, error) {
    return userRoleRepo.GetRolesByUserID(userID)
}

func RemoveUserRole(userRoleID uint) error {
    return userRoleRepo.Delete(userRoleID)
}

func CreateRole(name string) (*model.Role, error) {
    existing, _ := roleRepo.GetByName(name)
    if existing != nil {
        return nil, errors.New("role already exists")
    }
    role := &model.Role{Name: name}
    if err := roleRepo.Create(role); err != nil {
        return nil, err
    }
    return role, nil
}
