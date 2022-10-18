package user

import "time"

type User struct {
	ID       uint64    `json:"id" gorm:"column:id;primary_key" validate:"required"`
	Username string    `json:"username" gorm:"column:username"`
	Email    string    `json:"email" gorm:"column:email"`
	Password string    `json:"password" gorm:"column:password"`
	DOB      string    `json:"dob" gorm:"column:dob"`
	Age      uint64    `json:"age" gorm:"column:age"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at;default:CURRENT_TIMESTAMP"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at;omitempty"`
	DeleteAt time.Time `json:"delete_at" gorm:"column:delete_at;omitempty"`
}
