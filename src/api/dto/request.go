// src/api/dto/request.go
package dto

// RegisterRequest represents user registration request
// @Description User registration information
type RegisterRequest struct {
    Username     string `json:"username" binding:"required,min=5,max=20" example:"john_doe"`
    Password     string `json:"password" binding:"required,password,min=8" example:"Password123"`
    Email        string `json:"email" binding:"required,email" example:"john@example.com"`
    FirstName    string `json:"firstName" binding:"required,min=2,max=50" example:"John"`
    LastName     string `json:"lastName" binding:"required,min=2,max=50" example:"Doe"`
    MobileNumber string `json:"mobileNumber" binding:"required,mobile" example:"09123456789"`
}