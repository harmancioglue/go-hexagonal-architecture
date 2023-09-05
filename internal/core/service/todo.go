package service

import (
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/request"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/response"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/port/repository"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/port/service"
)

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) service.TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (tS *todoService) SaveTask(request *request.SaveTodoTaskRequest) *response.SaveTodoTaskResponse {
	var res *response.SaveTodoTaskResponse

	defer func() {
		if r := recover(); r != nil {
			res = &response.SaveTodoTaskResponse{
				Status:    false,
				ErrorCode: 0,
			}
		}
	}()

	todoEntity := tS.todoRepository.Save(request)

	res = &response.SaveTodoTaskResponse{
		Status: true,
		Data:   todoEntity,
	}

	return res
}
