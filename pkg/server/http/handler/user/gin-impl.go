package user

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/helper"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/message"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
	"github.com/gin-gonic/gin"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsecase}
}

func (u *UserHdlImpl) RegisterUserHdl(ctx *gin.Context) {
	log.Printf("%T - RegisterUserHdl is invoked]\n", u)
	defer log.Printf("%T - RegisterUserHdl executed\n", u)

	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed to bind payload",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PAYLOAD",
				"error_message": err.Error(),
			},
		})
		return
	}

	currAge := helper.AgeCalculator(time.Time(user.DOB), time.Now())
	user.Age = currAge
	if currAge < 8 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Your age must >= 8 years old",
		})
		return
	}

	log.Println("calling register user usecase")
	result, errMsg := u.userUsecase.RegisterUserSvc(ctx, user)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "User created successfully",
		"type":    "ACCEPTED",
		"data": gin.H{
			"age":      currAge,
			"email":    result.Email,
			"id":       result.ID,
			"username": result.Username,
		},
	})
}

func (u *UserHdlImpl) GetUserByIDHdl(ctx *gin.Context) {
	log.Printf("%T - GetUserByIDHdl is invoked]\n", u)
	defer log.Printf("%T - GetUserByIDHdl executed\n", u)

	userIdParam := ctx.Param("user_id")

	userId, err := strconv.ParseUint(userIdParam, 0, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 96,
			"type": "BAD_REQUEST",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PAYLOAD",
				"error_message": err.Error(),
			},
		})
		return
	}

	log.Println("calling get user by id usecase")
	result, errMsg := u.userUsecase.GetUserByIdSvc(ctx, userId)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "User Founded",
		"type":    "SUCCESS",
		"data": gin.H{
			"id":            result.ID,
			"username":      result.Username,
			"social_medias": result.SocialMedias,
		},
	})
}

func (u *UserHdlImpl) UpdateUserHdl(ctx *gin.Context) {
	log.Printf("%T - UpdateUserHdl is invoked\n", u)
	defer log.Printf("%T - UpdateUserHdl executed\n", u)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	var updatedUser user.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed to bind payload",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PAYLOAD",
				"error_message": err.Error(),
			},
		})
		return
	}

	idToken, errMsg := u.userUsecase.UpdateUserSvc(ctx, userId, updatedUser.Email, updatedUser.Username)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "user has been successfully updated",
		"type":    "SUCCESS",
		"data": gin.H{
			"id_token": idToken,
		},
	})
}

func (u *UserHdlImpl) DeleteUserHdl(ctx *gin.Context) {
	log.Printf("%T - DeleteUserHdl is invoked\n", u)
	defer log.Printf("%T - DeleteUserHdl executed\n", u)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	errMsg := u.userUsecase.DeleteUserSvc(ctx, userId)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "Your account has been successfully deleted",
		"type":    "SUCCESS",
	})
}
