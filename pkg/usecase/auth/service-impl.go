package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/helper"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/auth"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/claim"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/message"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/usecase/crypto"
	"github.com/google/uuid"
)

type AuthUsecaseImpl struct {
	authRepo    auth.AuthRepo
	userUsecase user.UserUsecase
}

func NewAuthUsecase(authRepo auth.AuthRepo, userUsecase user.UserUsecase) auth.AuthUsecase {
	return &AuthUsecaseImpl{authRepo: authRepo, userUsecase: userUsecase}
}

func (a *AuthUsecaseImpl) LoginUserSvc(ctx context.Context, email string, password string) (accessToken string, refreshToken string, idToken string, errMsg message.ErrorMessage) {
	log.Printf("%T - LoginUserSvc is invoked\n", a)
	defer log.Printf("%T - LoginUserSvc executed\n", a)

	log.Println("calling login user repo")
	result, err := a.authRepo.LoginUser(ctx, email)

	if result.ID <= 0 {
		err = errors.New("user not found")
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "USER_NOT_FOUND",
		}
		return accessToken, refreshToken, idToken, errMsg
	}

	if err != nil {
		log.Printf("error when fetching data from database: %s\n", err.Error())
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "INTERNAL_CONNECTION_PROBLEM",
		}
		return accessToken, refreshToken, idToken, errMsg
	}

	comparePass := helper.CheckPasswordHash(
		password, result.Password, 
	)

	if !comparePass {
		err := errors.New("invalid email or password")
		errMsg = message.ErrorMessage{
			Error: err,
			Type:  "WRONG_PASSWORD",
		}
		return accessToken, refreshToken, idToken, errMsg
	}

	timeNow := time.Now()
	claimAccess := claim.JWTToken{
		JWTID:          uuid.New(),
		Subject:        fmt.Sprintf("%v", result.ID),
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com",
		Scope:          "create update delete read",
		Type:           "ACCESS_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),
	}

	accessToken, _ = crypto.CreateJWT(ctx, claimAccess)

	claimRefresh := claim.JWTToken{
		JWTID:          uuid.New(),
		Subject:        fmt.Sprintf("%v", result.ID),
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com",
		Scope:          "create update delete read",
		Type:           "REFRESH_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(1000 * time.Hour).Unix(),
	}
	refreshToken, _ = crypto.CreateJWT(ctx, claimRefresh)

	claimId := claim.IDToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      time.Time(result.DOB),
	}
	idToken, _ = crypto.CreateJWT(ctx, claimId)

	return accessToken, refreshToken, idToken, errMsg
}

func (a *AuthUsecaseImpl) GetRefreshTokenSvc(ctx context.Context) (accessToken string, refreshToken string, idToken string, errMsg message.ErrorMessage) {
	log.Printf("%T - GetRefreshTokenSvc is invoked\n", a)
	defer log.Printf("%T - GetRefreshTokenSvc executed\n", a)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	log.Println("calling get user by id usecase")
	result, errMsg := a.userUsecase.GetUserByIdSvc(ctx, userId)

	if errMsg.Error != nil {
		return accessToken, refreshToken, idToken, errMsg
	}

	timeNow := time.Now()
	claimAccess := claim.JWTToken{
		JWTID:          uuid.New(),
		Subject:        stringUserId,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com",
		Scope:          "create update delete read",
		Type:           "ACCESS_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),
	}

	accessToken, _ = crypto.CreateJWT(ctx, claimAccess)

	claimRefresh := claim.JWTToken{
		JWTID:          uuid.New(),
		Subject:        stringUserId,
		Issuer:         "mygram.com",
		Audience:       "user.mygram.com",
		Scope:          "create update delete read",
		Type:           "REFRESH_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(1000 * time.Hour).Unix(),
	}
	refreshToken, _ = crypto.CreateJWT(ctx, claimRefresh)

	claimId := claim.IDToken{
		JWTID:    uuid.New(),
		Username: result.Username,
		Email:    result.Email,
		DOB:      time.Time(result.DOB),
	}
	idToken, _ = crypto.CreateJWT(ctx, claimId)

	return accessToken, refreshToken, idToken, errMsg
}


