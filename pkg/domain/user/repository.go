package user

import "context"

type UserRepo interface {
	GetUserByID(ctx context.Context, id string) (user User, err error)
	InsertUser(ctx context.Context, user *User) (err error)
}
