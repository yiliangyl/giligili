package utils

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CommonResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ErrResponse(code int, errMsg string) Response {
	return Response{
		Code: code,
		Message: errMsg,
		Data: nil,
	}
}
