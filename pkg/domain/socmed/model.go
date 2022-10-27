package socmed

import "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/gormmodel"

type SocialMedia struct {
	gormmodel.GormModel
	Name   string `gorm:"not null" json:"name" valid:"required~social media name is required"`
	URL    string `gorm:"not null" json:"url" valid:"required~social media url is required,url~invalid url format"`
	UserID uint64 `json:"user_id"`
}

func (SocialMedia) TableName() string {
	return "social_medias"
}
