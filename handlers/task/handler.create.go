package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/helpers"
	"github.com/habibiiiberahim/go-task/models"
	services "github.com/habibiiiberahim/go-task/services/task"
)


type handlerCreate struct {
	service services.ServiceCreate
}

func NewHandlerCreateTask(service services.ServiceCreate) *handlerCreate  {
	return &handlerCreate{
		service: service,
	}
}

func (h *handlerCreate)CreateTaskHandler(c *fiber.Ctx) error  {
	var input models.ModelTask
	c.BodyParser(&input)

	res, err := h.service.CreateTaskService(&input)
	
	switch err.Type{
	case "error_01":
		resp := helpers.ApiResponse(c, err.Code, false, "Task is not found", nil)
		return c.Status(err.Code).JSON(resp)
	case "error_02":
		resp := helpers.ApiResponse(c, err.Code, false, "Task is can't create", nil)
		return c.Status(err.Code).JSON(resp)
	default:
		resp :=	helpers.ApiResponse(c, fiber.StatusOK, true, "Task Create successfully", res)
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}
