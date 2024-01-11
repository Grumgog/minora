package routes

import (
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

	if strings.HasPrefix(path, "/static") {
		file := strings.Replace(path, "/static", "/", 1)
		c.File("./static" + file)
	}

	if !strings.HasSuffix(path, ".html") {
		path = path + ".html"
	}
	c.HTML(http.StatusOK, path, getTemplateParameter(path))
	return
}

func getTemplateParameter(path string) map[string]string {
	return map[string]string{"hello": "world"}
}
