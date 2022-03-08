package routers

import (
	"github.com/anika308142/mongoapi/controllers"
	"github.com/gin-gonic/gin"
)

func MyRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/movies", controllers.ReadAllMovie)
	router.GET("/movies/:id", controllers.ReadOneMovie)
	router.POST("/movies", controllers.AddMovie)
	router.PATCH("/movies/:id", controllers.MarkAsWatched)
	router.DELETE("/movies/:id", controllers.RemoveOneMovie)
	router.Run("localhost:9090")
	return router
}
