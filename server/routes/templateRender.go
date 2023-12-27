package routes

import (
	"os"
	"path"

	"keeper/utils"

	"github.com/gin-gonic/gin"
)

func AddTemplateRenderPath(engine *gin.Engine) {
	engine.GET("/*any", renderTemplateProcessing)
}

func renderTemplateProcessing(c *gin.Context) {
	urlPath := c.Request.URL.Path
	ext := path.Ext(urlPath)

	switch ext {
	case ".html":
		//template := renderTemplate(urlPath)
		break
	case ".css":
	case ".js":
		data, err := os.ReadFile(urlPath)
		utils.HandleErrorWithInternalServerError(c, err)
		c.Writer.Write(data)
	}

}

func renderTemplate(urlPath string) string {
	return "UnImplemented"
}
