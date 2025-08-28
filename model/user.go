// D:\Programing\projects\user-crud\model\user.go
package model

type User struct {
	BaseModel
	Username     string      `gorm:"type:varchar(20);not null;unique"`
	Password     string      `gorm:"type:varchar(64);not null"`
	UserRoles    *[]UserRole
}
