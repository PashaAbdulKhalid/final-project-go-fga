package user

import (
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"time"

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

func (u *UserHdlImpl) GetUserByIDHdl(ctx *gin.Context) {
	log.Printf("%T - GetUserByIDHdl is invoked]\n", u)
	defer log.Printf("%T - GetUserByIDHdl executed\n", u)

	// get id from request
	id := ctx.Param("id")
	// call service/usecase untuk menginsert data
	log.Println("calling get service usecase")
	result, err := u.userUsecase.GetUserByIDSvc(ctx, id)
	if err != nil {
		switch err.Error() {
		case "NOT_FOUND":
			ctx.AbortWithStatusJSON(http.StatusNotFound, message.Response{
				Code:  80,
				Error: "user not found",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success get user by id",
		Data:    result,
	})
}

func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	log.Printf("%T - InsertUserHdl is invoked]\n", u)
	defer log.Printf("%T - InsertUserHdl executed\n", u)

	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")
	var user user.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "failed to bind payload",
		})
		return
	}
	// check apakah username kosong atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check username from request")
	if user.Username == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "username should not be empty",
		})
		return
	}
	// check apakah email kosong atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check email from request")
	if user.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "email should not be empty",
		})
		return
	}

	// check apakah format email benar atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check format email from request")
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "email should use valid format",
		})
		return
	}

	// check apakah password kosong atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check password from request")
	if user.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "password should not be empty",
		})
		return
	}

	// check apakah panjang password lebih dari char ? lempar BAD_REQUEST
	log.Println("check password from request")
	if len(user.Password) < 6 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "password should more than 5 characters",
		})
		return
	}

	// check apakah age kosong atau tidak: kalau kosong lempar BAD_REQUEST
	dateFormat := "2006-01-02"
	dob, _ := time.Parse(dateFormat, user.DOB)
	age := time.Now().Year() - dob.Year()
	user.Age = age
	log.Println("check age from request")
	if user.Age == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "age should not be empty",
		})
		return
	}
	if user.Age <= 8 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "age should more than 8 years old",
		})
		return
	}
	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")
	result, err := u.userUsecase.InsertUserSvc(ctx, user)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		case "DUPLICATE_DATA":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  81,
				Error: "username or email already exist",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success insert user",
		Data:    result,
	})
}

func (u *UserHdlImpl) UpdateUserHdl(ctx *gin.Context) {
	log.Printf("%T - UpdateUserHdl is invoked]\n", u)
	defer log.Printf("%T - UpdateUserHdl executed\n", u)

	// get id from params
	log.Println("get id from params")
	id := ctx.Param("id")
	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")
	var user user.User
	// convert uint64 to string
	Id, _ := strconv.ParseUint(id, 10, 64)
	user.ID = Id
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "failed to bind payload",
		})
		return
	}

	// check apakah format email benar atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check format email from request")
	if user.Email != "" {
		_, err := mail.ParseAddress(user.Email)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "email should use valid format",
			})
			return
		}
	}

	// check apakah panjang password lebih dari char ? lempar BAD_REQUEST
	log.Println("check password from request")
	if user.Password != "" {
		if len(user.Password) < 6 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "password should more than 5 characters",
			})
			return
		}
	}

	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")
	result, err := u.userUsecase.UpdateUserSvc(ctx, user)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		case "DUPLICATE_DATA":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  81,
				Error: "username or email already exist",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success Update user",
		Data:    result,
	})
	
}

func (u *UserHdlImpl) DeleteUserHdl(ctx *gin.Context) {
	log.Printf("%T - DeleteUserByHdl is invoked]\n", u)
	defer log.Printf("%T - DeleteUserByHdl executed\n", u)

	// get id from request
	id := ctx.Param("id")
	var user user.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "failed to bind payload",
		})
		return
	}

	user.ID, _ = strconv.ParseUint(id, 10, 64)
	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")
	result, err := u.userUsecase.DeleteUserSvc(ctx, user)
	if err != nil {
		switch err.Error() {
		case "NOT_FOUND":
			ctx.AbortWithStatusJSON(http.StatusNotFound, message.Response{
				Code:  80,
				Error: "user not found",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success delete user by id",
		Data:    result.DeleteAt,
	})
}
