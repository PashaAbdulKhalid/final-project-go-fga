package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PashaAbdulKhalid/final-project-go-fga/config/postgres"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/message"
	"github.com/gin-gonic/gin"

	engine "github.com/PashaAbdulKhalid/final-project-go-fga/config/gin"
	userrepo "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/repository/user"
	userhandler "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/handler/user"
    userrouter "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/server/http/router/v1"
	userusecase "github.com/PashaAbdulKhalid/final-project-go-fga/pkg/usecase/user"
)

func main() {
	postgresCln := postgres.NewPostgresConnection(postgres.Config{
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		Password:     "12345",
		DatabaseName: "final-project",
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
		respMap := map[string]any{
			"code":       0,
			"message":    "server up and running",
			"start_time": startTime,
		}

		var respStruct message.Response

		resByte, err := json.Marshal(respMap)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(resByte, &respStruct)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, respStruct)
	})

	userRepo := userrepo.NewUserRepo(postgresCln)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	useHandler := userhandler.NewUserHandler(userUsecase)
	userrouter.NewUserRouter(ginEngine, useHandler).Router()

	ginEngine.Serve()

}
