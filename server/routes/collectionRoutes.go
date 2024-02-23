package routes

import "github.com/gin-gonic/gin"

func AddCollectionRoutes(routeGroup *gin.RouterGroup) {
	collection := routeGroup.Group("/collection")
	collection.GET("/", getCollections)
	collection.POST("/", createCollection)
	collection.GET("/{id}", getCollectionById)
	collection.PATCH("/{id}", updateCollection)
}

// getCollections godoc
// @summary Возвращает список всех коллекций на сервере
// @description Возвращает список всех коллекций на сервере.
// @tags collection
// @produce aplication/json
// @success 200 {object} response.CollectionListResponse
// @Router /collection/ [get]
func getCollections(c *gin.Context) {
}

// getCollection godoc
// @summary Возвращает коллекцию на сервере.
// @description Возвращает коллекцию на сервере, по id.
// @tags collection
// @Param        id   path      int  true  "Идентификатор коллекции"
// @produce aplication/json
// @success 200 {object} response.GetCollectionResponse
// @Router /collection/{id} [get]
func getCollectionById(c* gin.Context) {

}

func findCollection(c* gin.Context) {

}

// createCollections godoc
// @summary Создаёт коллекцию на сервере.
// @description Создаёт новую коллекцию на сервере.
// @tags collection
// @param request body request.CreateCollectionRequest true "Параметры создания коллекции"
// @produce aplication/json
// @success 201
// @Router /collection/ [post]
func createCollection(c *gin.Context) {

}


func updateCollection(c *gin.Context)

