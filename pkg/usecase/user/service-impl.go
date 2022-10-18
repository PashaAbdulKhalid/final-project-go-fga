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
	return result, err
}

func (u *UserUsecaseImpl) InsertUserSvc(ctx context.Context, input user.User) (result user.User, err error) {
	log.Printf("%T - InsertUserSvc is invoked]\n", u)
	defer log.Printf("%T - InsertUserSvc executed\n", u)

	log.Println("Inserting user to database")
	if err = u.userRepo.InsertUser(ctx, &input); err != nil {
		log.Printf("error when inserting user:%v\n", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
	}
	return input, err
}
