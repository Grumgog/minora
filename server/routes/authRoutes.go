package routes

import (
	"keeper/data/request"
	"keeper/data/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(routeGroup *gin.RouterGroup) {
	auth := routeGroup.Group("/auth")
	auth.POST("/login", Login)
}

// Login godoc
// @summary Авторизирует пользователя.
// @description Проводит авторизацию на сервере и возвращает информацию о пользователе.
// @tags auth
// @produce aplication/json
// @param request body request.LoginRequest true "Параметры аунтификации"
// @success 200 {object} response.LoginInfo
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var login request.LoginRequest
	if c.ShouldBindJSON(&login) != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var data = response.LoginInfo{
		Id:           0,
		Username:     "Admin",
		Email:        "admin@admin",
		CreatedAt:    time.Now(),
		LastLogin:    time.Now(),
		LastActionAt: time.Now(),
		LastAction:   response.Action(response.READ),
	}
	c.JSON(http.StatusOK, data)
}
