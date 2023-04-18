package routers

import (
	"simple_rest_api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:BookId", controllers.UpdateBook)

	router.GET("/books/:BookId", controllers.GetBook)

	// router.DELETE("/cars/:CarId", controllers.DeleteCar)

	return router
}
