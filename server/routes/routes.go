package routes

import "github.com/gin-gonic/gin"

func AddRoutes(engine *gin.Engine) {
	AddAuthRoutes(engine)
}
