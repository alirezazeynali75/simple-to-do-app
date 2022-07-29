package interfaces

import "github.com/gin-gonic/gin"

type Api interface {
	RegisterRouters(r *gin.Engine) *gin.Engine
}