package config

import (
	"time"
)

type Config struct {
	Redis      RedisConfig
	OTP        OTPConfig
	JWT        JWTConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type OTPConfig struct {
	Digits       int
	ExpireTime   time.Duration
	MaxAttempts  int
}

type JWTConfig struct {
	Secret           string
	RefreshSecret    string
	AccessTokenExpireTime  time.Duration
	RefreshTokenExpireTime time.Duration
}

func LoadConfig() *Config {
	return &Config{
		Redis: RedisConfig{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		OTP: OTPConfig{
			Digits:      6,
			ExpireTime:  2 * time.Minute,
			MaxAttempts: 3,
		},
		JWT: JWTConfig{
			Secret:          "your-secret-key",
			RefreshSecret:   "your-refresh-secret-key",
			AccessTokenExpireTime:  15 * time.Minute,
			RefreshTokenExpireTime: 7 * 24 * time.Hour,
		},
	}
}