package middlewares

import (
	"net/http"
	"strconv"

	"github.com/alirezazeynali75/simple-to-do-app/core/interfaces"
	"github.com/alirezazeynali75/simple-to-do-app/presentation/response_model"
	"github.com/gin-gonic/gin"
)

type Authentication struct {
	JwtManager interfaces.JwtManager
}

func (auth *Authentication) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header["Authorization"]
		claim, err := auth.JwtManager.Verify(token[0])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, response_model.NotOkResponse{
				BaseResponse: response_model.BaseResponse{
					Status: false,
				},
				Error: response_model.ErrorDetail{
					Code: response_model.FALIED_AUTH,
					Message: err.Error(),
				},
			})
		}
		ctx.Request.Header.Add("USER_ID", strconv.Itoa(int(claim.Id)))
		ctx.Next()
	}
}