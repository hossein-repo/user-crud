package usecase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"user-crud/config"
	"user-crud/model"
	"user-crud/repository"

	redisClient "user-crud/infra/redis" // نام مستعار
)

type OTPUsecase struct {
	cfg     *config.OTPConfig
	otpRepo repository.OTPRepository
}

func NewOTPUsecase(cfg *config.Config, otpRepo repository.OTPRepository) *OTPUsecase {
	return &OTPUsecase{
		cfg:     &cfg.OTP,
		otpRepo: otpRepo,
	}
}

func (u *OTPUsecase) GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	min := int64(100000)
	max := int64(999999)
	return strconv.FormatInt(rand.Int63n(max-min+1)+min, 10)
}

func (u *OTPUsecase) SendOTP(mobileNumber string) error {
	existingOTP, err := u.otpRepo.GetByMobileNumber(mobileNumber)
	if err == nil && existingOTP != nil {
		if time.Until(existingOTP.ExpiresAt) > time.Duration(u.cfg.ExpireTime/2) {
			return errors.New("active OTP already exists")
		}
	}

	otpCode := u.GenerateOTP()
	expiresAt := time.Now().Add(u.cfg.ExpireTime)

	ctx := context.Background()
	key := fmt.Sprintf("otp:%s", mobileNumber)

	err = redisClient.Client.Set(ctx, key, otpCode, u.cfg.ExpireTime).Err()
	if err != nil {
		return fmt.Errorf("failed to store OTP in redis: %w", err)
	}

	otp := &model.OTP{
		MobileNumber: mobileNumber,
		Code:         otpCode,
		ExpiresAt:    expiresAt,
		Used:         false,
		Attempts:     0,
	}

	if err := u.otpRepo.Create(otp); err != nil {
		redisClient.Client.Del(ctx, key)
		return fmt.Errorf("failed to create OTP record: %w", err)
	}

	fmt.Printf("OTP for %s: %s (Expires: %v)\n", mobileNumber, otpCode, expiresAt)
	return nil
}

func (u *OTPUsecase) VerifyOTP(mobileNumber, code string) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("otp:%s", mobileNumber)

	storedCode, err := redisClient.Client.Get(ctx, key).Result()
	if err == nil {
		if storedCode == code {
			redisClient.Client.Del(ctx, key)
			otp, err := u.otpRepo.GetByMobileNumber(mobileNumber)
			if err == nil && otp != nil {
				otp.Used = true
				u.otpRepo.Update(otp)
			}
			return true, nil
		}
	}

	otp, err := u.otpRepo.GetByMobileNumber(mobileNumber)
	if err != nil {
		return false, errors.New("OTP not found or expired")
	}

	if otp.Attempts >= u.cfg.MaxAttempts {
		return false, errors.New("maximum attempts exceeded")
	}

	if otp.Code != code {
		otp.Attempts++
		u.otpRepo.Update(otp)
		return false, errors.New("invalid OTP code")
	}

	if time.Now().After(otp.ExpiresAt) {
		return false, errors.New("OTP expired")
	}

	otp.Used = true
	u.otpRepo.Update(otp)
	redisClient.Client.Set(ctx, key, code, time.Until(otp.ExpiresAt))

	return true, nil
}

// این متد را اضافه کنید - برای تمیزکاری خودکار OTPهای منقضی
func (u *OTPUsecase) CleanupExpiredOTPs() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		u.otpRepo.DeleteExpiredOTPs()
	}
}