package repository

import (
	"errors"
	"user-crud/infra/db"
	"user-crud/model"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetByID(id uint) (*model.Role, error)
	GetByName(name string) (*model.Role, error)
	Create(role *model.Role) error
}

type roleRepository struct{}

func NewRoleRepository() RoleRepository {
	return &roleRepository{}
}

func (r *roleRepository) Create(role *model.Role) error {
	return db.DB.Create(role).Error
}

func (r *roleRepository) GetByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := db.DB.First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetByName(name string) (*model.Role, error) {
	var role model.Role
	if err := db.DB.Where("name = ?", name).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &role, nil
}
