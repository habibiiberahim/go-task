package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/helpers"
	"github.com/habibiiiberahim/go-task/models"
	services "github.com/habibiiiberahim/go-task/services/task"
)


type handlerUpdate struct {
	service services.ServiceUpdate
}

func NewHandlerUpdateTask(service services.ServiceUpdate) *handlerUpdate  {
	return &handlerUpdate{
		service: service,
	}
}

func (h *handlerUpdate)UpdateTaskHandler(c *fiber.Ctx) error  {
	var input models.ModelTask
	id, _ := c.ParamsInt("id")
	input.ID = uint(id)
	c.BodyParser(&input)

	res, err := h.service.UpdateTaskService(&input)
	switch err.Type{
	case "error_01":
		resp :=helpers.ApiResponse(c, err.Code, false, "Task is not found", nil)
		return c.Status(err.Code).JSON(resp)
	case "error_02":
		resp := helpers.ApiResponse(c, err.Code, false, "Task is can't update", nil)
		return c.Status(err.Code).JSON(resp)
	default:
		resp := helpers.ApiResponse(c, fiber.StatusOK, true, "Task update successfully", res)
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}