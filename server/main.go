package main

import (
	"fmt"
	"net/http"

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
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.AddRoutes(r)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("we can ping!")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/sample", SampleData)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Sample struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Lorem string `json:"lorem"`
}

// Sample data	godoc
// @summary return sample data from server
// @description return sample data as json.
// @produce application/json
// @success 200 {object} Sample
// @Router /sample [get]
func SampleData(c *gin.Context) {
	var data = Sample{
		Id:    1,
		Login: "admin",
		Lorem: "Ipsum",
	}
	c.JSON(http.StatusOK, data)
}
