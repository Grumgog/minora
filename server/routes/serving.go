package routes

import (
	"keeper/dbcontext"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddServingRoutes(engine *gin.Engine) {
	engine.LoadHTMLGlob("./templates/*")
	engine.GET("/*path", handlerForServing)
}

func handlerForServing(c *gin.Context) {
	path := c.Param("path")
	if path == "/" {
		c.HTML(http.StatusOK, "index.html", getTemplateParameter(path))
		return
	}

	if path == "/favicon.ico" {
		c.File("./favicon.ico")
	}

	if strings.HasPrefix(path, "/static") {
		file := strings.Replace(path, "/static", "/", 1)
		c.File("./static" + file)
		return
	}

	if !strings.HasSuffix(path, ".html") {
		path = path + ".html"
	}
	c.HTML(http.StatusOK, path, getTemplateParameter(path))
	return
}

type Parameters = map[string]string

func getTemplateParameter(path string) Parameters {
	templateParameter := Parameters{}
	getSystemParameter(&templateParameter)
	getParameterForTemplate(&templateParameter, path)
	return templateParameter
}

func getSystemParameter(parameters *Parameters) {
	db := dbcontext.GetDBContext()
	for key, value := range db.ParameterValue.GetSystemParameterValues() {
		(*parameters)[key] = value
	}
}

func getParameterForTemplate(parameters *Parameters, template string) {

}
