package response_model

type OkResponse struct {
	BaseResponse
	Message string `json:"message"`
	Info PaginationResponse `json:"info"`
	Result interface{}	`json:"result"`
}