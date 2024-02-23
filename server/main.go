package main

import (
	"minora/dbcontext"
	_ "minora/docs"
	"minora/routes"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title keeper cms server
// @version 0.1
// @description CMS с сервером на языке программирования GO

// @host localhost:8080
// @BasePath /api
func main() {
	dbcontext.Connect()
	app := fiber.New()
	app.Get("/docs/*any", fiberSwagger.WrapHandler)
	routes.AddRoutes(app)
	app.Listen(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
