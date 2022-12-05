package utils

const SuccessMsg = "success"

type BaseResponse struct {
	Message string `json:"message"`
	Data    interface{}
}

func SuccessResponse(data interface{}) BaseResponse {
	return BaseResponse{
		Message: SuccessMsg,
		Data:    data,
	}
}
