// api/dto/user.go
package dto

type RegisterUserRequest struct {
    FirstName    string `json:"firstName" binding:"required,min=3"`
    LastName     string `json:"lastName" binding:"required,min=6"`
    Username     string `json:"username" binding:"required,min=5"`
    Email        string `json:"email" binding:"min=6,email"`
    MobileNumber string `json:"mobileNumber" binding:"mobile,min=11,max=11"`
    Password     string `json:"password" binding:"required,password,min=6"`
}

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
    ExpiresIn    int64  `json:"expiresIn"`
}