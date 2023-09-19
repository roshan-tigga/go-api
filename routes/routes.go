package routes

import (
	"crud/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/persons", controllers.GetPersons)
	r.GET("/persons/:id", controllers.GetPerson)
	r.POST("/persons", controllers.CreatePerson)
	r.PUT("/persons/:id", controllers.UpdatePerson)
	r.DELETE("/persons/:id", controllers.DeletePerson)

	return r
}
