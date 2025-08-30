package model

type Role struct {
    BaseModel
    Name  string `gorm:"type:varchar(20);not null;unique"`
    Users []UserRole
}
