package model

type UserRole struct {
    BaseModel
    UserID uint
    RoleID uint
    Role   Role
}
