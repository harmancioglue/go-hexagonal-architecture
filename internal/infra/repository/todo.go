package repository

import (
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/entity"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/request"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/port/repository"
	"time"
)

type todoRepository struct {
	database repository.Database
}

func NewTodoRepository(database repository.Database) repository.TodoRepository {
	return &todoRepository{
		database: database,
	}
}

func (tR *todoRepository) Save(req *request.SaveTodoTaskRequest) *entity.Todo {
	db := tR.database.GetDatabase()

	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		panic(err)
	}

	todoEntity := entity.Todo{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
		Completed:   false,
	}

	result := db.Create(&todoEntity)

	dbErr := result.Error
	if dbErr != nil {
		panic(err)
	}

	return &todoEntity
}
