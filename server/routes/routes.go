package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	api := app.Group("/api")
	AddAuthRoutes(api)
	AddCollectionRoutes(api)
}
