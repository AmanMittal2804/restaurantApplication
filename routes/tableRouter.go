package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurantmanagement/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:invoice_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:invoice_id", controller.UpdateTable())
}
