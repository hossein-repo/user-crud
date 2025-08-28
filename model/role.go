// D:\Programing\projects\user-crud\model\role.go
package model

type Role struct {
	BaseModel
	Name  string      `gorm:"type:string;size:20;not null;unique"`
	Users *[]UserRole
}
