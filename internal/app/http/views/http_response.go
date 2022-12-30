package views

import "github.com/gofiber/fiber/v2"

type HttpResponse struct {
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
}

type ResponseMessage = map[string]string

func RenderResponse(c *fiber.Ctx, data interface{}) error {
	c.Status(fiber.StatusOK)
	return c.JSON(HttpResponse{Data: data})
}

func RenderErrorResponse(c *fiber.Ctx, status int, errorMsg string) error {
	c.Status(status)
	return c.JSON(HttpResponse{ErrorMessage: errorMsg})
}
