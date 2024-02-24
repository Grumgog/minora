package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleErrorWithInternalServerError(c *fiber.Ctx, err error) {
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
}

func ResponseWithError(c *fiber.Ctx, statusCode int, err error) {
	if err != nil {
		c.Status(statusCode).SendString(err.Error())
	}
}

func HandleErrorWithPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
