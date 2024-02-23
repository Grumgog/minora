package main

import (
	"minora/dbcontext"
	_ "minora/docs"
	"minora/routes"

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
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.AddRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
