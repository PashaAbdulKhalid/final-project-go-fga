package sosmed

import (
	"time"

	"gorm.io/gorm"
)

type Sosmed struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Url       string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}