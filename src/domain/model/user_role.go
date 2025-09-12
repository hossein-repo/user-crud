// src/domain/model/user_role.go
package model

type UserRole struct {
    BaseModel
    UserID uint `gorm:"not null" json:"userId"`
    RoleID uint `gorm:"not null" json:"roleId"`
    User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
    Role   Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"role,omitempty"`
}