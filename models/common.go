package models

import (
	"database/sql"
	"time"
)

// DeletedAt DeletedAt
// type DeletedAt sql.NullTime

// Common 公有字段
type Common struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"create_at" gorm:"create_at"`
	UpdatedAt time.Time    `json:"update_at" gorm:"update_at"`
	DeletedAt sql.NullTime `json:"delete_at" gorm:"index"`
}
