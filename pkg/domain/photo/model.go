package photo

import (
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/comment"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/gormmodel"
)

type Photo struct {
	gormmodel.GormModel
	Title    string            `gorm:"not null" json:"title" valid:"required~photo title is required"`
	Caption  string            `json:"caption"`
	Url      string            `gorm:"not null" json:"url" valid:"required~photo url is required,url~invalid url format"`
	UserID   uint64            `json:"user_id"`
	Comments []comment.Comment `json:"comments"`
}
