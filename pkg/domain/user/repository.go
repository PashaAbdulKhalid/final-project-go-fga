package user

import "context"

type UserRepo interface {
	GetUserByID(ctx context.Context, id string) (user User, err error)
	InsertUser(ctx context.Context, user *User) (err error)
	GetUserByUsername(ctx context.Context, username string) (user User, err error)
	GetUserByEmail(ctx context.Context, email string) (user User, err error)
	UpdateUser(ctx context.Context, user *User) (err error)
	DeleteUser(ctx context.Context, user *User) (err error)
}
