// src/domain/model/user.go
package model

type User struct {
    BaseModel
    Username     string     `gorm:"type:varchar(20);not null;unique" json:"username"`
    FirstName    string     `gorm:"type:varchar(15);null" json:"firstName"`
    LastName     string     `gorm:"type:varchar(25);null" json:"lastName"`
    MobileNumber string     `gorm:"type:varchar(11);null;unique" json:"mobileNumber"`
    Email        string     `gorm:"type:varchar(64);null;unique" json:"email"`
    Password     string     `gorm:"type:varchar(255);not null" json:"-"`
    Enabled      bool       `gorm:"default:true" json:"enabled"`
    UserRoles    []UserRole `gorm:"foreignKey:UserID" json:"userRoles,omitempty"`
}