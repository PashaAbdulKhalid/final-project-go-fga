package v1

import (
	engine "github.com/PashaAbdulKhalid/final-project-go-fga/config/gin"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/mygram/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) post() {
	u.routerGroup.POST("", u.userHandler.InsertUserHdl)
}

func (u *UserRouterImpl) Router() {
	u.post()
}
