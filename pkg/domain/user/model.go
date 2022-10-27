package user

import (
	"encoding/json"
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/helper"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/comment"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/gormmodel"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/photo"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/socmed"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type cusTime datatypes.Date

var _ json.Unmarshaler = &cusTime{}

func (ct *cusTime) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*ct = cusTime(t)

	return nil
}

type User struct {
	gormmodel.GormModel
	Username     string               `gorm:"not null;uniqueIndex" json:"username" valid:"required~username is required"`
	Email        string               `gorm:"not null;uniqueIndex" json:"email" valid:"required~email is required,email~invalid email format"`
	Password     string               `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~password must be at least 6 characters long"`
	DOB          cusTime              `gorm:"not null" json:"dob" valid:"required~Date of birth is required"`
	Age          int                  `gorm:"not null" json:"age"`
	SocialMedias []socmed.SocialMedia `json:"social_medias"`
	Photos       []photo.Photo        `json:"photos"`
	Comments     []comment.Comment    `json:"comments"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helper.HashPassword(u.Password)
	return
}
