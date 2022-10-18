package user

import (
	"log"
	"net/http"

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
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success insert user",
		Data:    result,
	})
}