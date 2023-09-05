package response

import (
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/entity"
)

type SaveTodoTaskResponse struct {
	Status       bool         `json:"status"`
	ErrorCode    int32        `json:"errorCode"`
	ErrorMessage string       `json:"errorMessage"`
	Data         *entity.Todo `json:"todoTask"`
}
