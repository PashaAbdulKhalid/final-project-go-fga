package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/claim"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/message"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/usecase/crypto"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) RegisterUserSvc(ctx context.Context, user user.User) (result user.User, errMsg message.ErrorMessage) {
	log.Printf("%T - InsertUserSvc is invoked]\n", u)
	defer log.Printf("%T - InsertUserSvc executed\n", u)

	// input validation
	if isValid, err := govalidator.ValidateStruct(user); !isValid {
		switch err.Error() {
		case "username is required":
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "USERNAME_IS_EMPTY",
			}
			return result, errMsg
		case "email is required":
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "EMAIL_IS_EMPTY",
			}
			return result, errMsg
		case "invalid email format":
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "INVALID_EMAIL_FORMAT",
			}
			return result, errMsg
		case "password is required":
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "PASSWORD_IS_EMPTY",
			}
			return result, errMsg
		case "password must be at least 6 characters long":
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "INVALID_PASSWORD_FORMAT",
			}
			return result, errMsg

		default:
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "INVALID_PAYLOAD",
			}
			return result, errMsg
		}
	}

	log.Println("calling register user repo")
	err := u.userRepo.RegisterUser(ctx, &user)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "idx_username"`) {
			err = errors.New("username has already been registered")
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "USERNAME_REGISTERED",
			}
			return result, errMsg
		}

		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "idx_email"`) {
			err = errors.New("email has already been registered")
			errMsg = message.ErrorMessage{
				Error: err,
				Type:  "EMAIL_REGISTERED",
			}
			return result, errMsg
		}
	}

	if err != nil {
		log.Printf("error when fetching data from database:%s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return result, errMsg

	}
	return user, errMsg
}

func (u *UserUsecaseImpl) GetUserByIdSvc(ctx context.Context, userId uint64) (result user.User, errMsg message.ErrorMessage) {
	log.Printf("%T - GetUserByID is invoked]\n", u)
	defer log.Printf("%T - GetUserByID executed\n", u)
	// get user from repository (database)
	log.Println("getting user from user repository")
	result, err := u.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Println("error when fetching data from database:  %s\n" + err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return result, errMsg
	}
	// check user id > 0 ?
	log.Println("checking user id")
	if result.ID <= 0 {
		// kalau tidak berarti user not found
		// log.Println("user with id %v is not found", userId)
		err = fmt.Errorf("user with id %v is not found", userId)
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "USER_NOT_FOUND",
		}
		return result, errMsg
	}
	return result, errMsg
}

func (u *UserUsecaseImpl) UpdateUserSvc(ctx context.Context, userId uint64, email string, username string) (idToken string, errMsg message.ErrorMessage) {
	log.Printf("%T - DeleteUserSvc is invoked]\n", u)
	defer log.Printf("%T - DeleteUserSvc executed\n", u)
	// check email validation
	if email == "" {
		errMsg := message.ErrorMessage{
			Error: errors.New("email is required"),
			Type:  "EMAIL_IS_EMPTY",
		}
		return idToken, errMsg
	}

	if !govalidator.IsEmail(email) {
		errMsg := message.ErrorMessage{
			Error: errors.New("invalid email format"),
			Type:  "INVALID_EMAIL_FORMAT",
		}
		return idToken, errMsg
	}

	// check username validation
	if username == "" {
		errMsg := message.ErrorMessage{
			Error: errors.New("username is required"),
			Type:  "USERNAME_IS_EMPTY",
		}
		return idToken, errMsg
	}

	// update user
	result, err := u.userRepo.UpdateUser(ctx, userId, email, username)

	if err != nil {
		log.Println("error when fetching data from database:  %s\n" + err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return idToken, errMsg
	}

	claimId := claim.IDToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      time.Time(result.DOB),
	}
	idToken, _ = crypto.CreateJWT(ctx, claimId)

	return idToken, errMsg
}

func (u *UserUsecaseImpl) DeleteUserSvc(ctx context.Context, userId uint64) (errMsg message.ErrorMessage) {
	log.Printf("%T - DeleteUserSvc is invoked]\n", u)
	defer log.Printf("%T - DeleteUserSvc executed\n", u)
	// delete user
	err := u.userRepo.DeleteUser(ctx, userId)

	if err != nil {
		log.Println("error when fetching data from database:  %s\n" + err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return errMsg
	}

	return errMsg
}
