package repository

import (
	"time"
	"user-crud/infra/db"
	"user-crud/model"
)

type OTPRepository interface {
	Create(otp *model.OTP) error
	GetByMobileNumber(mobileNumber string) (*model.OTP, error)
	Update(otp *model.OTP) error
	DeleteExpiredOTPs() error
}

type otpRepo struct{}

func NewOTPRepository() OTPRepository {
	return &otpRepo{}
}

func (r *otpRepo) Create(otp *model.OTP) error {
	return db.DB.Create(otp).Error
}

func (r *otpRepo) GetByMobileNumber(mobileNumber string) (*model.OTP, error) {
	var otp model.OTP
	err := db.DB.Where("mobile_number = ? AND used = ? AND expires_at > ?", 
		mobileNumber, false, time.Now()).First(&otp).Error
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *otpRepo) Update(otp *model.OTP) error {
	return db.DB.Save(otp).Error
}

func (r *otpRepo) DeleteExpiredOTPs() error {
	return db.DB.Where("expires_at < ?", time.Now()).Delete(&model.OTP{}).Error
}