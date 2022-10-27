package user

import (
	"context"
	"log"

	"github.com/PashaAbdulKhalid/final-project-go-fga/config/postgres"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
	"gorm.io/gorm/clause"
)

type UserRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewUserRepo(pgCln postgres.PostgresClient) user.UserRepo {
	return &UserRepoImpl{pgCln: pgCln}
}

func (u *UserRepoImpl) RegisterUser(ctx context.Context, insertedUser *user.User) (err error) {
	log.Printf("%T - InsertUser is invoked]\n", u)
	defer log.Printf("%T - InsertUser executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()

	// insert new user
	err = db.Model(&user.User{}).Create(&insertedUser).Error
	//check error
	if err != nil {
		log.Printf("error when inserting user to database ")
		return err
	}
	return err
}

func (u *UserRepoImpl) GetUserByID(ctx context.Context, userId uint64) (result user.User, err error) {
	log.Printf("%T - GetUserByID is invoked]\n", u)
	defer log.Printf("%T - GetUserByID executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// get user
	err = db.Model(&user.User{}).Where("id = ?", userId).Preload("SocialMedias").Find(&result).Error
	//check error
	if err != nil {
		log.Printf("error when getting user with ID %v\n",
			userId)
	}
	return result, err
}

func (u *UserRepoImpl) UpdateUser(ctx context.Context, userId uint64, email string, username string) (result user.User, err error) {
	log.Printf("%T - UpdateUser is invoked]\n", u)
	defer log.Printf("%T - UpdateUser executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// update user
	err = db.Model(&result).Clauses(clause.Returning{}).Where("id = ?", userId).Updates(user.User{Email: email, Username: username}).Error
	//check error
	if err != nil {
		log.Printf("error when updating user with ID %v\n",
			userId)
	}
	return result, err
}

func (u *UserRepoImpl) DeleteUser(ctx context.Context, userId uint64) (err error) {
	log.Printf("%T - DeleteUser is invoked]\n", u)
	defer log.Printf("%T - DeleteUser executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()

	// deletedUser.DeleteAt = time.Now()
	// delete user
	err = db.Where("id = ?", userId).Delete(&user.User{}).Error
	//check error
	if err != nil {
		log.Printf("error when deleting user with ID %v\n",
			userId)
	}
	return err
}
