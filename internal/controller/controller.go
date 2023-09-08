package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/request"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/model/response"
	"github.com/harmancioglue/go-hexagonal-architecture/internal/core/port/service"
)

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(serv service.TodoService) *todoController {
	return &todoController{
		todoService: serv,
	}
}

func (tc *todoController) Run() {
	app := fiber.New()
	app.Post("/todo-task", func(ctx *fiber.Ctx) error {
		var reqObj request.SaveTodoTaskRequest
		err := ctx.BodyParser(&reqObj)
		if err != nil {
			return ctx.JSON(fiber.Map{
				"Response": response.SaveTodoTaskResponse{
					Status:       false,
					ErrorCode:    0,
					ErrorMessage: "Check request body",
					Data:         nil,
				},
			})
		}

		resp := tc.todoService.SaveTask(&reqObj)
		return ctx.JSON(fiber.Map{
			"Response": resp,
		})
	})

	app.Listen(":4000")
}
