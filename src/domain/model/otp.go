// src/domain/model/otp.go
package model

import "time"

type OTP struct {
    BaseModel
    MobileNumber string    `gorm:"type:varchar(11);not null" json:"mobileNumber"`
    Code         string    `gorm:"type:varchar(6);not null" json:"code"`
    ExpiresAt    time.Time `gorm:"not null" json:"expiresAt"`
    Used         bool      `gorm:"default:false" json:"used"`
    Attempts     int       `gorm:"default:0" json:"attempts"`
}