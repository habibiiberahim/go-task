package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/helpers"
	"github.com/habibiiiberahim/go-task/models"
	services "github.com/habibiiiberahim/go-task/services/task"
)


type handlerDone struct {
	service services.ServiceDone
}

func NewHandlerDoneTask(service services.ServiceDone) *handlerDone  {
	return &handlerDone{
		service: service,
	}
}

func (h *handlerDone)DoneTaskHandler(c *fiber.Ctx) error  {
	var input models.ModelTask
	id, _ := c.ParamsInt("id")
	input.ID = uint(id)

	res, err := h.service.DoneTaskService(&input)
	
	switch err.Type{
	case "error_01":
		resp :=helpers.ApiResponse(c, err.Code, false, "Task is not found", nil)
		return c.Status(err.Code).JSON(resp)
	case "error_02":
		resp :=helpers.ApiResponse(c, err.Code, false, "Task is can't done", nil)
		return c.Status(err.Code).JSON(resp)
	default:
		resp := helpers.ApiResponse(c, fiber.StatusOK, true, "Task done successfully", res)
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}