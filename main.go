package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alirezazeynali75/simple-to-do-app/core/interfaces"
	"github.com/alirezazeynali75/simple-to-do-app/core/services"
	"github.com/alirezazeynali75/simple-to-do-app/data/database/mysql"
	"github.com/alirezazeynali75/simple-to-do-app/data/database/mysql/model"
	"github.com/alirezazeynali75/simple-to-do-app/data/repos"
	"github.com/alirezazeynali75/simple-to-do-app/presentation/api"
	"github.com/alirezazeynali75/simple-to-do-app/utils"
	"github.com/gin-gonic/gin"
)

func parent() {
	db := mysql.OpenConnection()
	defer mysql.KillConnection(db)
	r := gin.Default()
	userApi := api.UserApi{
		Us: &services.UserService{
			Repo: &repos.UserRepo{
				BaseRepos: repos.BaseRepos{
					Db: db,
				},
			},
			JwtManager: *interfaces.GetJwtManagerInstance("Alirez@1375", time.Duration(time.Hour * 2)),
		},
	}
	engine := userApi.RegisterRoutes(r)
	engine.Run()
}

func main() {
	utils.LoadEnv()
	db := mysql.OpenConnection()
	defer mysql.KillConnection(db)
	err := db.AutoMigrate(
		&model.Status{},
		&model.Tasks{},
		&model.User{},
	)
	if err != nil {
		panic(err)
	}
	go parent()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM)
	fmt.Println("press ctrl + c to exit program")
	<- ch
}