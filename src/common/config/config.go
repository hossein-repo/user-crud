// src/common/config/config.go
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config ساختار اصلی پیکربندی برنامه
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Password PasswordConfig `mapstructure:"password"`
	Otp      OtpConfig      `mapstructure:"otp"`
}

// Load بارگذاری تنظیمات از فایل و محیط
func Load() (*Config, error) {
	viper.SetConfigName("config")    // نام فایل: config.yaml
	viper.SetConfigType("yaml")      // نوع فایل: yaml
	viper.AddConfigPath(".")         // جستجو در دایرکتوری جاری
	viper.AddConfigPath("./config")  // جستجو در پوشه config
	viper.AddConfigPath("../config") // جستجو در پوشه والد/config

	// خواندن از environment variables
	viper.AutomaticEnv()

	// تنظیم مقادیر پیش‌فرض
	setDefaults()

	// خواندن فایل config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

// setDefaults مقادیر پیش‌فرض برای تنظیمات
func setDefaults() {
	viper.SetDefault("app.name", "user-crud")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.env", "development")

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.timeout", 30*time.Second)

	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.timezone", "Asia/Tehran")
	viper.SetDefault("database.max_open_conns", 25)
	viper.SetDefault("database.max_idle_conns", 5)

	viper.SetDefault("jwt.access_token_expire", 15*time.Minute)
	viper.SetDefault("jwt.refresh_token_expire", 168*time.Hour) // 7 days

	viper.SetDefault("password.min_length", 8)
	viper.SetDefault("password.require_uppercase", true)
	viper.SetDefault("password.require_lowercase", true)
	viper.SetDefault("password.require_digits", true)

	viper.SetDefault("otp.digits", 6)
	viper.SetDefault("otp.expire_time", 120) // seconds
}
