// src/common/config/jwt.go
package config

import "time"

type JWTConfig struct {
    Secret                     string        `mapstructure:"secret"`
    RefreshSecret              string        `mapstructure:"refresh_secret"`
    AccessTokenExpireDuration  time.Duration `mapstructure:"access_token_expire"`
    RefreshTokenExpireDuration time.Duration `mapstructure:"refresh_token_expire"`
    Issuer                     string        `mapstructure:"issuer"`
}