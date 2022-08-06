package api

import (
	"net/http"

	"github.com/alirezazeynali75/simple-to-do-app/presentation/dto"
	"github.com/alirezazeynali75/simple-to-do-app/presentation/ports"
	"github.com/alirezazeynali75/simple-to-do-app/presentation/response_model"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	Us ports.UserService
}

func (ua *UserApi) SignUp(ctx *gin.Context) {
	dto := &dto.UserSignUpDto{}
	err := ctx.BindJSON(dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response_model.NotOkResponse{
			BaseResponse: response_model.BaseResponse{
				Status: false,
			},
			Error: response_model.ErrorDetail{
				Message: err.Error(),
				Code: response_model.VALIDATION_FAILED,
			},
		})
	} else {
		user, err := ua.Us.SignUp(dto.Username, dto.Email, dto.PhoneNumber, dto.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response_model.NotOkResponse{
				BaseResponse: response_model.BaseResponse{
					Status: false,
				},
				Error: response_model.ErrorDetail{
					Message: err.Error(),
					Code: response_model.SIGN_UP_FAILED,
				},
			})
		} else {
			ctx.JSON(http.StatusOK, response_model.OkResponse{
				BaseResponse: response_model.BaseResponse{
					Status: true,
				},
				Message: "operation was successful",
				Info: response_model.PaginationResponse{
					Page: 1,
				},
				Result: user,
			})
		}
	}
}

func (ua *UserApi) Login(ctx *gin.Context) {
	dto := &dto.UserLoginDto{}
	err := ctx.BindJSON(dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response_model.NotOkResponse{
			BaseResponse: response_model.BaseResponse{
				Status: false,
			},
			Error: response_model.ErrorDetail{
				Message: err.Error(),
				Code: response_model.VALIDATION_FAILED,
			},
		})
	} else {
		bearerToken, user, err := ua.Us.Login(dto.Email, dto.Password)
		type result struct {
			User interface{} `json:"user"`
			Token string `json:"token"`
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response_model.NotOkResponse{
				BaseResponse: response_model.BaseResponse{
					Status: false,
				},
				Error: response_model.ErrorDetail{
					Message: err.Error(),
					Code: response_model.SIGN_UP_FAILED,
				},
			})
		} else {
			ctx.JSON(http.StatusOK, response_model.OkResponse{
				BaseResponse: response_model.BaseResponse{
					Status: true,
				},
				Message: "operation was successful",
				Info: response_model.PaginationResponse{
					Page: 1,
				},
				Result: result{
					User: user,
					Token: bearerToken,
				},
			})
		}
	}
}

func (ua *UserApi) RegisterRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/sign_up", ua.SignUp)
	r.POST("/login", ua.Login)
	return r
}