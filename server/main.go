package main

import (
	"minora/dbcontext"
	_ "minora/docs"
	"minora/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title MinoraCMS server
// @version 0.1
// @description CMS с сервером на языке программирования GO

// @host localhost:8080
// @BasePath /api
func main() {
	dbcontext.Connect()
	apiApp := fiber.New()
	apiApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	apiApp.Get("/docs/*", fiberSwagger.WrapHandler)
	routes.AddRoutes(apiApp)
	go apiApp.Listen("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	Apps := fiber.New()
	Apps.Static("/cms", "/cms")
	Apps.Static("/", "/site")
	Apps.Listen("localhost:80")	
}
