package routes

import (
	"minora/handler/collection_handler"
	"minora/server/server_errors"
	"minora/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddCollectionRoutes(routeGroup fiber.Router) {
	collection := routeGroup.Group("/collection")
	collection.Get("/", getCollections)
	collection.Post("/", createCollection)
	collection.Get("/:id", getCollectionById)
	collection.Patch("/:id", updateCollection) 
}

// getCollections godoc
// @summary Возвращает список всех коллекций на сервере
// @description Возвращает список всех коллекций на сервере.
// @tags collection
// @produce aplication/json
// @success 200 {object} response.CollectionListResponse
// @Router /collection/ [get]
func getCollections(c *fiber.Ctx) error {
	result := collection_handler.GetList()
	return c.JSON(result)
}

// getCollection godoc
// @summary Возвращает коллекцию на сервере.
// @description Возвращает коллекцию на сервере, по id.
// @tags collection
// @Param id path int true "Идентификатор коллекции"
// @produce aplication/json
// @success 200 {object} response.GetCollectionResponse
// @success 404
// @Router /collection/{id} [get]
func getCollectionById(c *fiber.Ctx) error {
	id, error := c.ParamsInt("id")
	utils.ResponseWithError(c, http.StatusBadRequest, error)
	result, err := collection_handler.GetById(uint(id))
	if err != nil {
		errInfo := server_errors.HandleServerError(c, err)
		return c.Status(errInfo.StatusCode).SendString(errInfo.Error.Error())
	}
	return c.JSON(result)
}

// createCollections godoc
// @summary Создаёт коллекцию на сервере.
// @description Создаёт новую коллекцию на сервере.
// @tags collection
// @param request body request.CreateCollectionRequest true "Параметры создания коллекции"
// @produce aplication/json
// @success 201
// @Router /collection/ [post]
func createCollection(c *fiber.Ctx) error {
	return nil
}

// updateCollection godoc
// @summary Редактирует коллекцию.
// @description Редактирует информацию о коллекции на сервере.
// @tags collection
// @param request body request.UpdateCollectionRequest true "Параметры создания коллекции"
// @produce aplication/json
// @success 200
// @Router /collection/ [patch]
func updateCollection(c *fiber.Ctx) error {
	return nil
}

