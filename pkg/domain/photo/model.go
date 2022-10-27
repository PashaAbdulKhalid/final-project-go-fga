package photo

import (
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/comment"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	Url       string            `gorm:"not null"`
	UserID    uint              `gorm:"not null"`
	Comment   []comment.Comment `gorm:"foreignKey:PhotoID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
