package routes

import (
	"shwetaik-sql-acc-backend-api/controllers"

	"github.com/labstack/echo/v4"
)

func PaymentRoutes(e *echo.Group, controller *controllers.PaymentController) {
	e.GET("/payments", controller.GetAll)
	e.GET("/payments/:docKey", controller.GetByDOCKEY)
	e.POST("/payments", controller.Create)
}

func PaymentMethodRoutes(e *echo.Group, controller *controllers.PaymentMethodController) {
	e.GET("/paymentMethods", controller.GetAll)
	e.GET("/paymentMethods/:code", controller.GetByCode)
}

func PaymentDetailRoutes(e *echo.Group, controller *controllers.PaymentDetailController) {
	e.GET("/paymentDetails", controller.GetAll)
}
