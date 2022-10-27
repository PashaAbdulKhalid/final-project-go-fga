package auth

import (
	"context"
	"log"

	"github.com/PashaAbdulKhalid/final-project-go-fga/config/postgres"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/auth"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
)

type AuthRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewAuthRepo(pgCln postgres.PostgresClient) auth.AuthRepo {
	return &AuthRepoImpl{pgCln: pgCln}
}

func (a *AuthRepoImpl) LoginUser(ctx context.Context, email string) (result user.User, err error) {
	log.Printf("%T - LoginUser is invoked\n", a)
	defer log.Printf("%T - LoginUser executed\n", a)

	db := a.pgCln.GetClient()

	err = db.Model(&user.User{}).Select("id", "username", "password", "email", "dob").Where("email = ?", email).Find(&result).Error

	if err != nil {
		log.Printf("error when getting email %v\n", email)
	}

	return result, err
}
