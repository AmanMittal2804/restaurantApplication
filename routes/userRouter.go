package routes

import (
	controller "restaurantmanagement/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/login", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
