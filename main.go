package main

import (
	"net/http"
	"os"
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/config/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	engine "github.com/PashaAbdulKhalid/final-project-go-fga/config/gin"
	userrepo "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/repository/user"
	userhandler "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/handler/user"
	userusecase "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/usecase/user"

	authrepo "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/repository/auth"
	authhandler "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/handler/auth"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/middleware"
	authusecase "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/usecase/auth"

	router "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/router/v1"
)

func init() {
	godotenv.Load(".env")

}
func main() {
	postgresHost := os.Getenv("MY_GRAM_POSTGRES_HOST")
	postgresPort := os.Getenv("MY_GRAM_POSTGRES_PORT")
	postgresDatabase := os.Getenv("MY_GRAM_POSTGRES_DATABASE")
	postgresUsername := os.Getenv("MY_GRAM_POSTGRES_USERNAME")
	postgresPassword := os.Getenv("MY_GRAM_POSTGRES_PASSWORD")
	// sharedKey := os.Getenv("MY_GRAM_JWT_SHARED_KEY")

	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         postgresHost,
		Port:         postgresPort,
		User:         postgresUsername,
		Password:     postgresPassword,
		DatabaseName: postgresDatabase,
	})

	// gin engine
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})
	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "server up and running",
			"started": startTime,
		})
	})

	userRepo := userrepo.NewUserRepo(postgresCln)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	useHandler := userhandler.NewUserHandler(userUsecase)

	authRepo := authrepo.NewAuthRepo(postgresCln)
	authUsecase := authusecase.NewAuthUsecase(authRepo, userUsecase)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	authMiddleware := middleware.NewAuthMiddleware(userUsecase)

	router.NewUserRouter(ginEngine, useHandler, authMiddleware).Routers()
	router.NewAuthRouter(ginEngine, authHandler, authMiddleware).Routers()

	ginEngine.Serve()

}
