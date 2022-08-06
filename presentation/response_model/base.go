package response_model


type BaseResponse struct {
	Status bool `json:"status"`
}

type PaginationResponse struct {
	Page uint	`json:"page"`
}

