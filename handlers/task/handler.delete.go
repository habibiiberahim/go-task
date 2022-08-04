package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/helpers"
	"github.com/habibiiiberahim/go-task/models"
	services "github.com/habibiiiberahim/go-task/services/task"
)


type handlerDelete struct {
	service services.ServiceDelete
}

func NewHandlerDeleteTask(service services.ServiceDelete) *handlerDelete  {
	return &handlerDelete{
		service: service,
	}
}

func (h *handlerDelete)DeleteTaskHandler(c *fiber.Ctx) error  {
	var input models.ModelTask
	id, _ := c.ParamsInt("id")
	input.ID = uint(id)

	res, err := h.service.DeleteTaskService(&input)
	
	switch err.Type{
	case "error_01":
		resp :=helpers.ApiResponse(c, err.Code, false, "Task is not found", nil)
		return c.Status(err.Code).JSON(resp)
	case "error_02":
		resp :=helpers.ApiResponse(c, err.Code, false, "Task is can't delete", nil)
		return c.Status(err.Code).JSON(resp)
	default:
		resp := helpers.ApiResponse(c, fiber.StatusOK, true, "Task delete successfully", res)
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}