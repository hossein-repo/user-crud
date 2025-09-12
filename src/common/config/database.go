// src/common/config/database.go
package config

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    SSLMode  string `mapstructure:"sslmode"`
    Timezone string `mapstructure:"timezone"`
    
    // Connection pool settings
    MaxOpenConns int `mapstructure:"max_open_conns"`
    MaxIdleConns int `mapstructure:"max_idle_conns"`
    MaxLifetime  int `mapstructure:"max_lifetime"` // in minutes
}

func (d *DatabaseConfig) DSN() string {
    return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
        d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode, d.Timezone)
}