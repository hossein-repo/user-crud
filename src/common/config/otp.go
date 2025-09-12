// src/common/config/otp.go
package config

import "time"

// OtpConfig تنظیمات سیستم OTP
type OtpConfig struct {
    Digits          int           `mapstructure:"digits"`
    ExpireTime      time.Duration `mapstructure:"expire_time"`
    MaxAttempts     int           `mapstructure:"max_attempts"`
    ResendDelay     time.Duration `mapstructure:"resend_delay"`
    BlockDuration   time.Duration `mapstructure:"block_duration"`
    AllowedChars    string        `mapstructure:"allowed_chars"`
}

// Validate بررسی اعتبار تنظیمات OTP
func (o *OtpConfig) Validate() error {
    if o.Digits < 4 || o.Digits > 8 {
        return fmt.Errorf("OTP digits must be between 4 and 8")
    }
    if o.ExpireTime < 30*time.Second {
        return fmt.Errorf("OTP expire_time must be at least 30 seconds")
    }
    return nil
}

// GetExpireTime مدت زمان انقضای OTP
func (o *OtpConfig) GetExpireTime() time.Duration {
    return o.ExpireTime * time.Second
}