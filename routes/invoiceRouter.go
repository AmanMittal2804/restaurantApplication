package routes

import (
	"github.com/gin-gonic/gin"

	controller "restaurantmanagement/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.POST("/invoices/:invoice_id", controller.UpdateInvoice())
}
