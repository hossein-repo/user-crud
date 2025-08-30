package model

import (
    "time"

    "gorm.io/gorm"
)

type BaseModel struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time      `gorm:"autoCreateTime"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime"`
    CreatedBy uint           `gorm:"not null;default:1"`
    UpdatedBy uint           `gorm:"not null;default:1"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
    if b.CreatedBy == 0 {
        b.CreatedBy = 1
    }
    if b.UpdatedBy == 0 {
        b.UpdatedBy = 1
    }
    return
}
