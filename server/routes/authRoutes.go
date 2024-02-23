package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddAuthRoutes(routeGroup fiber.Router) {
	auth := routeGroup.Group("/api")
	auth.Post("/login", Login)
	auth.Post("/change-password", ChangePassword)
}

// Login godoc
// @summary Авторизирует пользователя.
// @description Проводит авторизацию на сервере и возвращает информацию о пользователе.
// @tags auth
// @produce aplication/json
// @param request body request.LoginRequest true "Параметры аунтификации"
// @success 200 {object} response.LoginInfo
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	return nil
	// var login request.LoginRequest
	// if c.ShouldBindJSON(&login) != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	// response, err := handler.HandleAuth(login)
	// if err != nil {
	// 	utils.HandleErrorWithInternalServerError(c, err)
	// 	return
	// }

	// c.JSON(http.StatusOK, response)
}

// ChangePassword godoc
// @summary Меняет пароль для пользователя
// @description Меняет пароль по запросу пользователя
// @tags auth
// @produce aplication/json
// @param request body request.LoginRequest true "Параметры смены пароля"
// @success 200 {object} response.LoginInfo
// @Router /auth/change-password [post]
func ChangePassword(c *fiber.Ctx) error {
	return nil

}
