package v1

import (
	engine "github.com/PashaAbdulKhalid/final-project-go-fga/config/gin"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/user"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/middleware"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/router"
	"github.com/gin-gonic/gin"
)

type UserRouterImpl struct {
	ginEngine      engine.HttpServer
	routerGroup    *gin.RouterGroup
	userHandler    user.UserHandler
	authMiddleware middleware.AuthMiddleware
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler, authMiddleware middleware.AuthMiddleware) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler, authMiddleware: authMiddleware}
}

func (u *UserRouterImpl) post() {
	u.routerGroup.POST("/register", u.userHandler.RegisterUserHdl)
}

func (u *UserRouterImpl) get() {
	u.routerGroup.GET("/:user_id", u.userHandler.GetUserByIDHdl)
}

func (u *UserRouterImpl) put() {
	u.routerGroup.PUT("/:user_id", u.authMiddleware.CheckJWTAuth, u.userHandler.UpdateUserHdl)
}

func (u *UserRouterImpl) delete() {
	u.routerGroup.DELETE(":user_id", u.authMiddleware.CheckJWTAuth, u.userHandler.DeleteUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.post()
	u.get()
	u.put()
	u.delete()
}
