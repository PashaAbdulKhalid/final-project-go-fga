package auth

import (
	"context"

	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
)

type AuthRepo interface {
	LoginUser(ctx context.Context, email string) (result user.User, err error)
}