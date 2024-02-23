package routes

import "github.com/gin-gonic/gin"

func AddRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
	AddAuthRoutes(api)
	AddCollectionRoutes(api)
}
