// src/infra/db/migration.go
package db

import (
	"fmt"
	"math/rand"
	"time"
	
	"gorm.io/gorm"
	"user-crud/src/domain/model"
)

func MigrateExistingData() error {
	// ابتدا بررسی کنید که آیا نیاز به migration هست یا نه
	var nullCount int64
	err := DB.Model(&model.User{}).
		Where("first_name IS NULL OR last_name IS NULL OR mobile_number IS NULL").
		Count(&nullCount).Error
	
	if err != nil {
		return err
	}
	
	if nullCount == 0 {
		fmt.Println("✅ No NULL values found - skipping migration")
		return DB.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRole{})
	}
	
	fmt.Printf("🔧 Migrating %d records with NULL values\n", nullCount)
	
	// برای هر کاربر یک شماره موبایل منحصر به فرد تولید کنید
	err = DB.Transaction(func(tx *gorm.DB) error {
		// ابتدا کاربران با mobile_number NULL را پیدا کنید
		var users []model.User
		if err := tx.Where("mobile_number IS NULL").Find(&users).Error; err != nil {
			return err
		}
		
		// برای هر کاربر یک شماره موبایل منحصر به فرد تنظیم کنید
		for i, user := range users {
			user.FirstName = coalesce(user.FirstName, "Unknown")
			user.LastName = coalesce(user.LastName, "Unknown")
			user.MobileNumber = generateUniqueMobileNumber(i)
			
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
		}
		
		return nil
	})
	
	if err != nil {
		return err
	}
	
	fmt.Println("✅ Migration completed successfully")
	return DB.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRole{})
}

func coalesce(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func generateUniqueMobileNumber(index int) string {
	rand.Seed(time.Now().UnixNano() + int64(index))
	// تولید شماره موبایل منحصر به فرد
	return fmt.Sprintf("091%08d", rand.Intn(100000000))
}