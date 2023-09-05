package repository

import (
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/entity"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/request"
)

type TodoRepository interface {
	Save(req *request.SaveTodoTaskRequest) *entity.Todo
}
