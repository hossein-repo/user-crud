// src/common/config/password.go
package config

// PasswordConfig تنظیمات سیاست گذرواژه
type PasswordConfig struct {
    MinLength          int  `mapstructure:"min_length"`
    RequireUppercase   bool `mapstructure:"require_uppercase"`
    RequireLowercase   bool `mapstructure:"require_lowercase"`
    RequireDigits      bool `mapstructure:"require_digits"`
    RequireSpecialChars bool `mapstructure:"require_special_chars"`
    MaxAttempts        int  `mapstructure:"max_attempts"`
    LockoutTime        int  `mapstructure:"lockout_time"` // minutes
}

// Validate بررسی اعتبار تنظیمات گذرواژه
func (p *PasswordConfig) Validate() error {
    if p.MinLength < 6 {
        return fmt.Errorf("password min_length must be at least 6")
    }
    if p.MaxAttempts < 1 {
        return fmt.Errorf("max_attempts must be at least 1")
    }
    return nil
}