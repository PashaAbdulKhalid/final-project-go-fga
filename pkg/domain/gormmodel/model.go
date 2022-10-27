package gormmodel

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint64          `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
