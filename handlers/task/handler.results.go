package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/helpers"
	services "github.com/habibiiiberahim/go-task/services/task"
)


type handlerResults struct {
	service services.ServiceResults
}

func NewHandlerResultsTask(service services.ServiceResults) *handlerResults  {
	return &handlerResults{
		service: service,
	}
}

func (h *handlerResults)ResultsTaskHandler(c *fiber.Ctx) error  {
	res, err := h.service.ResultsTaskService()
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