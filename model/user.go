package model

type User struct {
    BaseModel
    Username     string      `gorm:"type:varchar(20);not null;unique"`
    Email        string      `gorm:"type:varchar(64);unique;default:null"`
    Password     string      `gorm:"type:varchar(64);not null"`
    Enabled      bool        `gorm:"default:true"`
    UserRoles    []UserRole
}
