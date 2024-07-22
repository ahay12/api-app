package helper

import "github.com/gofiber/fiber/v2"

type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
type ErrorField struct {
	ID      string `json:"id"`
	Value   string `json:"value"`
	Caused  string `json:"caused"`
	Message string `json:"message"`
}

func RespondJSON(ctx *fiber.Ctx, status int, message string, payload interface{}, errors interface{}) {
	var res ResponseData

	if status >= 200 && status < 300 {
		res.Success = true
	}

	if len(message) != 0 {
		res.Message = message
	}

	if payload != nil {
		res.Data = payload
	}

	if errors != nil {
		res.Error = errors
	}

	err := ctx.JSON(res)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
}
