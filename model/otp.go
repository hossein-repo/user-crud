
package model

import "time"

type OTP struct {
	BaseModel
	MobileNumber string    `gorm:"type:varchar(11);not null;uniqueIndex:idx_mobile_otp"`
	Code         string    `gorm:"type:varchar(6);not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	Used         bool      `gorm:"default:false"`
	Attempts     int       `gorm:"default:0"`
}