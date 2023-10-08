package pkg

import (
	"github.com/dhritimannath/notification-service/internal/constant"
	"github.com/dhritimannath/notification-service/internal/dtos"
)

func Null() interface{} {
 return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dtos.ApiResponse[T] {
 return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dtos.ApiResponse[T] {
 return dtos.ApiResponse[T]{
  ResponseKey:     status,
  ResponseMessage: message,
  Data:            data,
 }
}