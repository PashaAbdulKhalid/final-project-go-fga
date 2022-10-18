package user

import "context"

type UserUsecase interface {
	GetUserByIDSvc(ctx context.Context, id string) (result User, err error)
	InsertUserSvc(ctx context.Context, input User) (result User, err error)
	GetUserByUsernameSvc(ctx context.Context, username string) (result User, err error)
	GetUserByEmailSvc(ctx context.Context, email string) (result User, err error)
}
