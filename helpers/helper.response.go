package helpers

import "github.com/gofiber/fiber/v2"

func ApiResponse(c *fiber.Ctx, Code int, Success bool, Message string, Data interface{}) fiber.Map {

	resp := fiber.Map{
		"Code":    Code,
		"Success": Success,
		"Message": Message,
		"Data":    Data,
	}
	return resp
}