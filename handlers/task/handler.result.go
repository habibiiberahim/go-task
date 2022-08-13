package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/helpers"
	"github.com/habibiiiberahim/go-task/models"
	services "github.com/habibiiiberahim/go-task/services/task"
)


type handlerResult struct {
	service services.ServiceResult
}

func NewHandlerResultTask(service services.ServiceResult) *handlerResult  {
	return &handlerResult{
		service: service,
	}
}

func (h *handlerResult)ResultTaskHandler(c *fiber.Ctx) error  {
	var input models.ModelTask
	id, _ := c.ParamsInt("id")
	input.ID = uint(id)
	c.BodyParser(&input)
	
	res, err := h.service.ResultTaskService(&input)
	switch err.Type{
	case "error_01":
		resp :=helpers.ApiResponse(c, err.Code, false, "Task is not found", nil)
		return c.Status(err.Code).JSON(resp)
	case "error_02":
		resp :=helpers.ApiResponse(c, err.Code, false, "Task is can't read", nil)
		return c.Status(err.Code).JSON(resp)
	default:
		resp :=helpers.ApiResponse(c, fiber.StatusOK, true, "Task read successfully", res)
		return c.Status(fiber.StatusOK).JSON(resp)
	}
}