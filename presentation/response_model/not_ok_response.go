package response_model

type ErrorDetail struct {
	Code uint `json:"code"`
	Message string `json:"message"`
}

type NotOkResponse struct {
	BaseResponse
	Error ErrorDetail `json:"error_detail"`
}

const (
	VALIDATION_FAILED = 4001
	SIGN_UP_FAILED = 1001
	FALIED_AUTH = 4000
)