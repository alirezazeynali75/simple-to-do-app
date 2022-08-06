package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alirezazeynali75/simple-to-do-app/core/interfaces"
	"github.com/alirezazeynali75/simple-to-do-app/presentation/middlewares"
	"github.com/gin-gonic/gin"
)

type PingApi struct {
}

func (pa *PingApi) Ping(ctx *gin.Context) {
	fmt.Println(ctx.Request.Header.Get("USER_ID"))
	ctx.JSON(http.StatusAccepted, "")
}

func (ua *PingApi) RegisterRoutes(r *gin.Engine) *gin.Engine {
	auth := middlewares.Authentication{
		JwtManager: *interfaces.GetJwtManagerInstance("Alirez@1375", time.Duration(time.Hour * 2)),
	}
	authenticated := r.Group("/api/ping")
	authenticated.Use(auth.Auth())
	authenticated.GET("/", ua.Ping)
	return r
}
