// src/domain/model/user.go
package model

type User struct {
	BaseModel
	Username     string `gorm:"type:varchar(20);not null;unique"`
	FirstName    string `gorm:"type:varchar(50);not null"`
	LastName     string `gorm:"type:varchar(50);not null"`
	MobileNumber string `gorm:"type:varchar(11);unique;not null"`
	Email        string `gorm:"type:varchar(64);unique;not null"`
	Password     string `gorm:"type:varchar(128);not null"` // افزایش طول برای hash
	Enabled      bool   `gorm:"default:true"`
	UserRoles    []UserRole
}