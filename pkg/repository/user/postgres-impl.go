package user

import (
	"context"
	"log"

	"github.com/PashaAbdulKhalid/final-project-go-fga/config/postgres"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
)

type UserRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewUserRepo(pgCln postgres.PostgresClient) user.UserRepo {
	return &UserRepoImpl{pgCln: pgCln}
}

func (u *UserRepoImpl) GetUserByID(ctx context.Context, id string) (result user.User, err error) {
	return result, err
}

func (u *UserRepoImpl) InsertUser(ctx context.Context, insertedUser *user.User) (err error) {
	log.Printf("%T - InsertUser is invoked]\n", u)
	defer log.Printf("%T - InsertUser executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new user

	db.Model(&user.User{}).
		Create(&insertedUser)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when inserting user with email ")
	}
	return err
}
