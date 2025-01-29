package routes

import (
	"shwetaik-sql-acc-backend-api/controllers"
	"shwetaik-sql-acc-backend-api/middlewares"

	"github.com/labstack/echo/v4"
)

func PaymentRoutes(e *echo.Group, controller *controllers.PaymentController) {
	paymentRouteGroup := e.Group("/payments")
	paymentRouteGroup.Use(middlewares.AuthMiddleware)
	paymentRouteGroup.GET("", controller.GetAll)
	paymentRouteGroup.GET("/:docKey", controller.GetByDOCKEY)
	paymentRouteGroup.POST("", controller.Create)
}

func PaymentMethodRoutes(e *echo.Group, controller *controllers.PaymentMethodController) {
	e.GET("/paymentMethods", controller.GetAll)
	e.GET("/paymentMethods/:code", controller.GetByCode)
}

func PaymentDetailRoutes(e *echo.Group, controller *controllers.PaymentDetailController) {
	e.GET("/paymentDetails", controller.GetAll)
}
