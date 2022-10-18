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
	log.Printf("%T - GetUserByID is invoked]\n", u)
	defer log.Printf("%T - GetUserByID executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new user
	db.Model(&user.User{}).
		Where("id = ?", id).
		Find(&result)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when getting user with ID %v\n",
			id)
	}
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

func (u *UserRepoImpl) GetUserByUsername(ctx context.Context, username string) (result user.User, err error) {
	log.Printf("%T - GetUserByUsername is invoked]\n", u)
	defer log.Printf("%T - GetUserByUsername executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new user
	db.Model(&user.User{}).
		Where("username = ?", username).
		Find(&result)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when getting user with username %v\n",
			username)
	}
	return result, err
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (result user.User, err error) {
	log.Printf("%T - GetUserByEmail is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmail executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new user
	db.Model(&user.User{}).
		Where("email = ?", email).
		Find(&result)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when getting user with email %v\n",
			email)
	}
	return result, err
}
