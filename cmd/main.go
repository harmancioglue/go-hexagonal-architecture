package main

import (
	"github.com/harmancioglue/go-hexagonal-architecture/internal/controller"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/service"
	repository2 "github.com/harmancioglue/go-hexagonal-architecture/internal/infra/repository"
	"github.com/joho/godotenv"
)

func main() {
	//load env file
	godotenv.Load(".env")

	//create db
	db, err := repository2.NewDatabase()
	if err != nil {

	}

	//create repository
	todoRepository := repository2.NewTodoRepository(db)
	//create service
	todoService := service.NewTodoService(todoRepository)
	//create controller
	ctr := controller.NewTodoController(todoService)

	ctr.Run()

}
