package routes

import (
	"keeper/handler"
	"keeper/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddParameterRoutes(routeGroup *gin.RouterGroup) {
	parameter := routeGroup.Group("/parameter")
	parameter.GET("/", GetParameters)
}

// GetParameters godoc
// @summary Возвращает все параметры для использования в шаблонах.
// @description Возвращает массив параметров, предназначеных для использования в шаблонах.
// @tags parameter
// @produce aplication/json
// @success 200 {object} []response.Parameter
// @Router /parameter [get]
func GetParameters(c *gin.Context) {
	parameter, err := handler.ParameterGetListHandler()
	if err != nil {
		utils.HandleErrorWithInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, parameter)
}
