package repository

import (
    "user-crud/infra/db"
    "user-crud/model"
)

type RoleRepository interface {
    Create(role *model.Role) error
    GetByID(id uint) (*model.Role, error)
    GetByName(name string) (*model.Role, error)
}

type roleRepo struct{}

func NewRoleRepository() RoleRepository {
    return &roleRepo{}
}

func (r *roleRepo) Create(role *model.Role) error {
    return db.DB.Create(role).Error
}

func (r *roleRepo) GetByID(id uint) (*model.Role, error) {
    var role model.Role
    if err := db.DB.First(&role, id).Error; err != nil {
        return nil, err
    }
    return &role, nil
}

func (r *roleRepo) GetByName(name string) (*model.Role, error) {
    var role model.Role
    if err := db.DB.Where("name = ?", name).First(&role).Error; err != nil {
        return nil, err
    }
    return &role, nil
}
