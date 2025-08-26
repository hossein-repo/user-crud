package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	CreatedBy uint           `gorm:"not null;default:1"` // مقدار پیش‌فرض
	UpdatedBy uint           `gorm:"not null;default:1"` // مقدار پیش‌فرض
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate hook برای تمام مدل‌ها
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.CreatedBy == 0 {
		b.CreatedBy = 1 // یا هر ID پیش‌فرض دیگری که مدنظر است
	}
	if b.UpdatedBy == 0 {
		b.UpdatedBy = 1
	}
	return
}

// BeforeUpdate hook برای به‌روزرسانی UpdatedBy
func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// اگر میخوای می‌تونی اینجا هم مقدار UpdatedBy را خودکار تغییر بدی
	// b.UpdatedBy = currentUserID (مثلاً از context)
	return
}
