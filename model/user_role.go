// D:\Programing\projects\user-crud\model\user_role.go
package model

type UserRole struct {
	BaseModel
	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`
}
