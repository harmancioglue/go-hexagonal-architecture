package main

import (
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/controller"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/infra/repository"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/service"
	"github.com/joho/godotenv"
)

func main() {
	//load env file
	godotenv.Load(".env")

	//create db
	db, err := repository.NewDatabase()
	if err != nil {

	}

	//create repository
	todoRepository := repository.NewTodoRepository(db)
	//create service
	todoService := service.NewTodoService(todoRepository)
	//create controller
	ctr := controller.NewTodoController(todoService)

	ctr.Run()

}
