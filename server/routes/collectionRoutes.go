package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddCollectionRoutes(routeGroup fiber.Router) {
	collection := routeGroup.Group("/collection")
	collection.Get("/", getCollections)
	collection.Post("/", createCollection)
	collection.Get("/{id}", getCollectionById)
	collection.Patch("/{id}", updateCollection)
}

// getCollections godoc
// @summary Возвращает список всех коллекций на сервере
// @description Возвращает список всех коллекций на сервере.
// @tags collection
// @produce aplication/json
// @success 200 {object} response.CollectionListResponse
// @Router /collection/ [get]
func getCollections(c *fiber.Ctx) error {
	return nil
}

// getCollection godoc
// @summary Возвращает коллекцию на сервере.
// @description Возвращает коллекцию на сервере, по id.
// @tags collection
// @Param        id   path      int  true  "Идентификатор коллекции"
// @produce aplication/json
// @success 200 {object} response.GetCollectionResponse
// @Router /collection/{id} [get]
func getCollectionById(c *fiber.Ctx) error {
	return nil
}

func findCollection(c *fiber.Ctx) error {
	return nil
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


func updateCollection(c *fiber.Ctx) error {
	return nil
}

