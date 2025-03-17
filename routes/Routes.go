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
	paymentMethodRouteGroup := e.Group("/payment-methods", middlewares.AuthMiddleware)
	paymentMethodRouteGroup.GET("", controller.GetAll)
	paymentMethodRouteGroup.GET("/:code", controller.GetByCode)
}

func PaymentDetailRoutes(e *echo.Group, controller *controllers.PaymentDetailController) {
	paymentDetailRouteGroup := e.Group("/payment-details", middlewares.AuthMiddleware)
	paymentDetailRouteGroup.GET("", controller.GetAll)
}

func ProjectRoutes(e *echo.Group, controller *controllers.ProjectController) {
	projectRouteGroup := e.Group("/projects", middlewares.AuthMiddleware)
	projectRouteGroup.GET("", controller.GetAll)
	projectRouteGroup.GET("/:code", controller.GetByCode)
}

func GLAccRoutes(e *echo.Group, controller *controllers.GLAccController) {
	GLAccRouteGroup := e.Group("/gl-accounts")
	GLAccRouteGroup.GET("", controller.GetAll)
	GLAccRouteGroup.GET("/low-level", controller.GetAllLowLevel)
}
