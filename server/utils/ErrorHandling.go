package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
