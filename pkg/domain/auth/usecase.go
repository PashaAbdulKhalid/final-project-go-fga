package auth

import (
	"context"

	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/message"
)

type AuthUsecase interface {
	LoginUserSvc(ctx context.Context, email string, password string) (accessToken string, refreshToken string, idToken string, errMsg message.ErrorMessage)
	GetRefreshTokenSvc(ctx context.Context) (accessToken string, refreshToken string, idToken string, errMsg message.ErrorMessage)
}
