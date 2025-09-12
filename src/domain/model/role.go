// src/domain/model/role.go
package model

type Role struct {
    BaseModel
    Name      string     `gorm:"type:varchar(20);not null;unique" json:"name"`
    UserRoles []UserRole `gorm:"foreignKey:RoleID" json:"userRoles,omitempty"`
}