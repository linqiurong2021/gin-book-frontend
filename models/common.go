package models

import (
	"time"

	"gorm.io/gorm"
)

// Common 公有字段
type Common struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"create_at" gorm:"create_at;not null"`
	UpdatedAt time.Time      `json:"update_at" gorm:"update_at;not null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // 使用 gorm.DeletedAt 软删除
}
