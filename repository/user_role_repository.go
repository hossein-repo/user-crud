package repository

import (
    "errors"
    "user-crud/infra/db"
    "user-crud/model"
)

type UserRoleRepository interface {
    Create(userRole *model.UserRole) error
    GetRolesByUserID(userID uint) ([]model.Role, error)
    Delete(userRoleID uint) error
}

type userRoleRepo struct{}

func NewUserRoleRepository() UserRoleRepository {
    return &userRoleRepo{}
}

func (r *userRoleRepo) Create(userRole *model.UserRole) error {
    return db.DB.Create(userRole).Error
}
func (r *userRoleRepo) GetRolesByUserID(userID uint) ([]model.Role, error) {
    var userRoles []model.UserRole
    
    if err := db.DB.
        Preload("Role").
        Where("user_id = ?", userID).
        Find(&userRoles).Error; err != nil {
        return nil, err
    }
    
    // استفاده از map برای حذف نقش‌های تکراری
    uniqueRoles := make(map[uint]model.Role)
    for _, ur := range userRoles {
        if ur.Role.ID != 0 {
            uniqueRoles[ur.Role.ID] = ur.Role
        }
    }
    
    // تبدیل map به slice
    roles := make([]model.Role, 0, len(uniqueRoles))
    for _, role := range uniqueRoles {
        roles = append(roles, role)
    }
    
    return roles, nil
}
func (r *userRoleRepo) Delete(userRoleID uint) error {
    res := db.DB.Delete(&model.UserRole{}, userRoleID)
    if res.RowsAffected == 0 {
        return errors.New("userRole not found")
    }
    return res.Error
}
