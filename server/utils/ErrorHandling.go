package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrorWithInternalServerError(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func HandleErrorWithPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
