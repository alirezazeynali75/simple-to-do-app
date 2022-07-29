package main

import (
	"time"

	"github.com/alirezazeynali75/simple-to-do-app/data/database/repo"
	"github.com/alirezazeynali75/simple-to-do-app/domain"
	"github.com/alirezazeynali75/simple-to-do-app/domain/services"
	"github.com/alirezazeynali75/simple-to-do-app/presentation/api"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	api := api.UserApi{
		Service: services.UserService{
			UserRepo: &repo.UserRepo{},
			Jwt: domain.NewJWTManager("sss", time.Duration(1)),
		},
	}
}