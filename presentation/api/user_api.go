package api

import (
	"net/http"
	"time"

	"github.com/alirezazeynali75/simple-to-do-app/domain"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	Service *domain.UserService
}

func (api *UserApi) RegisterRouters(r *gin.Engine) *gin.Engine {
	type PingResponseModel struct {
		Status bool `json:"status"`
		Message string `json:"message"`
		Date time.Time `json:"_"`
	}
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, PingResponseModel{
			Status: true,
			Message: "PONG",
			Date: time.Now(),
		})
	})
	return r
}