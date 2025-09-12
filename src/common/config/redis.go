// src/common/config/redis.go
package config

import (
    "fmt"
    "strconv"
)

// RedisConfig تنظیمات اتصال به Redis
type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
    PoolSize int    `mapstructure:"pool_size"`
}

// Addr آدرس کامل Redis
func (r *RedisConfig) Addr() string {
    return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

// DSN رشته اتصال به Redis
func (r *RedisConfig) DSN() string {
    if r.Password != "" {
        return fmt.Sprintf("redis://:%s@%s:%d/%d", r.Password, r.Host, r.Port, r.DB)
    }
    return fmt.Sprintf("redis://%s:%d/%d", r.Host, r.Port, r.DB)
}

// IsEnabled بررسی فعال بودن Redis
func (r *RedisConfig) IsEnabled() bool {
    return r.Host != "" && r.Port > 0
}