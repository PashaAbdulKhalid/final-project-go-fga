package user

import (
	"context"
	"log"
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/config/postgres"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
	"golang.org/x/crypto/bcrypt"
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
		Where("id = ? AND delete_at < ?", id, "1000-01-01").
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
	// hashing password
	pass := []byte(insertedUser.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	insertedUser.Password = string(hashedPass)
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

func (u *UserRepoImpl) UpdateUser(ctx context.Context, updatedUser *user.User) (err error) {
	log.Printf("%T - UpdateUser is invoked]\n", u)
	defer log.Printf("%T - UpdateUser executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new user
	db.Model(&user.User{}).
		Where("id = ?", updatedUser.ID).
		Update("username", updatedUser.Username).
		Update("email", updatedUser.Email).
		Update("password", updatedUser.Password).
		Update("updated_at", time.Now())
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when updating user with ID %v\n",
			updatedUser.ID)
	}
	return err
}

func (u *UserRepoImpl) DeleteUser(ctx context.Context, deletedUser *user.User) (err error) {
	log.Printf("%T - DeleteUser is invoked]\n", u)
	defer log.Printf("%T - DeleteUser executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()

	deletedUser.DeleteAt = time.Now()
	// insert new user
	db.Model(&user.User{}).
		Where("id = ?", deletedUser.ID).
		Update("delete_at", deletedUser.DeleteAt)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when deleting user with ID %v\n",
			deletedUser.ID)
	}
	return err
}
