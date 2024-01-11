package main

import (
	"keeper/dbcontext"
	"keeper/routes"

	_ "keeper/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title keeper cms server
// @version 0.1
// @description CMS с сервером на языке программирования GO

// @host localhost:8080
// @BasePath /api
func main() {
	dbcontext.Connect()
	apiEngine := gin.Default()
	apiEngine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.AddRoutes(apiEngine)
	go apiEngine.Run(":8080")

	servingEngine := gin.Default()
	routes.AddServingRoutes(servingEngine)
	servingEngine.Run(":8081")
}
