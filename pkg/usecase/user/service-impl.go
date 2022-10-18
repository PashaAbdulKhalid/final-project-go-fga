package user

import (
	"context"
	"errors"
	"log"

	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) GetUserByIDSvc(ctx context.Context, id string) (result user.User, err error) {
	log.Printf("%T - GetUserByID is invoked]\n", u)
	defer log.Printf("%T - GetUserByID executed\n", u)
	// get user from repository (database)
	log.Println("getting user from user repository")
	result, err = u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Println("error when fetching data from database: " + err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
		return result, err
	}
	// check user id > 0 ?
	log.Println("checking user id")
	if result.ID <= 0 {
		// kalau tidak berarti user not found
		log.Println("user is not found: " + id)
		err = errors.New("NOT_FOUND")
		return result, err
	}
	return result, err
}

func (u *UserUsecaseImpl) InsertUserSvc(ctx context.Context, input user.User) (result user.User, err error) {
	log.Printf("%T - InsertUserSvc is invoked]\n", u)
	defer log.Printf("%T - InsertUserSvc executed\n", u)

	// check if username is already exist
	usrCheck1, err1 := u.GetUserByUsernameSvc(ctx, input.Username)
	if err1 == nil {
		// user found
		log.Printf("user has been registered with id: %v\n", usrCheck1.ID)
		err = errors.New("DUPLICATE_DATA")
		return result, err
	}
	// get user for input email first
	usrCheck, err := u.GetUserByEmailSvc(ctx, input.Email)

	// check user is exist or not
	if err == nil {
		// user found
		log.Printf("user has been registered with id: %v\n", usrCheck.ID)
		err = errors.New("DUPLICATE_DATA")
		return result, err
	}
	// internal server error condition
	if err.Error() != "NOT_FOUND" {
		// internal server error
		log.Println("got error when checking user from database")
		return result, err
	}

	log.Println("Inserting user to database")
	if err = u.userRepo.InsertUser(ctx, &input); err != nil {
		log.Printf("error when inserting user:%v\n", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
	}
	return input, err
}

func (u *UserUsecaseImpl) GetUserByUsernameSvc(ctx context.Context, username string) (result user.User, err error) {
	log.Printf("%T - GetUserByUsername is invoked]\n", u)
	defer log.Printf("%T - GetUserByUsername executed\n", u)
	// get user from repository (database)
	log.Println("getting user from user repository")
	result, err = u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Println("error when fetching data from database: " + err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
		return result, err
	}
	// check user id > 0 ?
	log.Println("checking user id")
	if result.ID <= 0 {
		// kalau tidak berarti user not found
		log.Println("user is not found: " + username)
		err = errors.New("NOT_FOUND")
		return result, err
	}
	return result, err
}

func (u *UserUsecaseImpl) GetUserByEmailSvc(ctx context.Context, email string) (result user.User, err error) {
	log.Printf("%T - GetUserByEmail is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmail executed\n", u)
	// get user from repository (database)
	log.Println("getting user from user repository")
	result, err = u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Println("error when fetching data from database: " + err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
		return result, err
	}
	// check user id > 0 ?
	log.Println("checking user id")
	if result.ID <= 0 {
		// kalau tidak berarti user not found
		log.Println("user is not found: " + email)
		err = errors.New("NOT_FOUND")
		return result, err
	}
	return result, err
}
