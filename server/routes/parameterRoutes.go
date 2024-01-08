package routes

import (
	"keeper/data/request"
	"keeper/handler"
	"keeper/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddParameterRoutes(routeGroup *gin.RouterGroup) {
	parameter := routeGroup.Group("/parameter")
	parameter.GET("/", GetParameters)
	parameter.POST("/new", CreateParameter)
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

// CreateParameter godoc
// @summary Создаёт параметр для редактирования в шаблоне.
// @description Создаёт параметр.
// @tags parameter
// @produce aplication/json
// @param request body request.CreateParameterRequest true "Свойства создаваемого шаблона"
// @success 200
// @Router /parameter/new [post]
func CreateParameter(c *gin.Context) {
	var requestData request.CreateParameterRequest
	c.BindJSON(&requestData)
	handler.ParameterCreateHandler(requestData.Name, requestData.ParameterType, requestData.IsDefault)
}
