package httpException

import (
	"learn-go/internal/constant"
	"learn-go/pkg/dto"
)

func Null() interface{} {
	return nil
}

// BuildResponse builds the API response with optional data
func BuildResponse(responseStatus constant.ResponseStatus, data ...interface{}) dto.ApiResponse[interface{}] {
	var message string
	if len(data) > 1 {
		message, _ = data[0].(string)
	} else {
		message = responseStatus.GetResponseMessage()
	}
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[len(data)-1]
	}
	return BuildResponse_(responseStatus.GetResponseStatus(), message, responseData)
}

// BuildResponse_ is the internal function to build the API response
func BuildResponse_(status string, message string, data interface{}) dto.ApiResponse[interface{}] {
	return dto.ApiResponse[interface{}]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
