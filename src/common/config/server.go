// src/common/config/server.go
package config

import "time"

type ServerConfig struct {
    Port         int           `mapstructure:"port"`
    Timeout      time.Duration `mapstructure:"timeout"`
    ReadTimeout  time.Duration `mapstructure:"read_timeout"`
    WriteTimeout time.Duration `mapstructure:"write_timeout"`
    Debug        bool          `mapstructure:"debug"`
}