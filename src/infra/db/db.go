// infra/db/db.go
package db

import (
    "fmt"
    "log"
    "time"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
    dsn := "host=localhost user=postgres password=admin dbname=car_sale_db port=5432 sslmode=disable TimeZone=Asia/Tehran"
    
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    
    if err != nil {
        return fmt.Errorf("failed to connect to database: %w", err)
    }
    
    // تنظیم connection pool
    sqlDB, err := DB.DB()
    if err != nil {
        return fmt.Errorf("failed to get database instance: %w", err)
    }
    
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)
    
    log.Println("✅ Database connected successfully to: car_sale_db")
    return nil
}

// AutoMigrate تابع جدید برای migration
func AutoMigrate(models ...interface{}) error {
    if DB == nil {
        return fmt.Errorf("database not initialized")
    }
    
    err := DB.AutoMigrate(models...)
    if err != nil {
        return fmt.Errorf("failed to auto migrate: %w", err)
    }
    
    log.Println("✅ Database migrations completed")
    return nil
}

// Close تابع برای بستن اتصال
func Close() error {
    if DB != nil {
        sqlDB, err := DB.DB()
        if err != nil {
            return err
        }
        return sqlDB.Close()
    }
    return nil
}