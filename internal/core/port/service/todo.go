package service

import (
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/request"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/response"
)

type TodoService interface {
	SaveTask(request *request.SaveTodoTaskRequest) *response.SaveTodoTaskResponse
}
