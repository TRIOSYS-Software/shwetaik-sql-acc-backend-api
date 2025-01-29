package main

import (
	"fmt"

	"shwetaik-sql-acc-backend-api/config"
	"shwetaik-sql-acc-backend-api/controllers"
	"shwetaik-sql-acc-backend-api/repositories"
	"shwetaik-sql-acc-backend-api/routes"
	"shwetaik-sql-acc-backend-api/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://yourdomain.com"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	cfg, err := config.GetConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}
	db, err := cfg.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	apiV1 := e.Group("/api/v1")

	paymentRepo := repositories.NewPaymentRepo(db)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService)
	routes.PaymentRoutes(apiV1, paymentController)

	paymentMethodRepo := repositories.NewPaymentMethodRepo(db)
	paymentMethodService := services.NewPaymentMethodService(paymentMethodRepo)
	paymentMethodController := controllers.NewPaymentMethodController(paymentMethodService)
	routes.PaymentMethodRoutes(apiV1, paymentMethodController)

	paymentDetailRepo := repositories.NewPaymentDetailRepo(db)
	paymentDetailService := services.NewPaymentDetailService(paymentDetailRepo)
	paymentDetailController := controllers.NewPaymentDetailController(paymentDetailService)
	routes.PaymentDetailRoutes(apiV1, paymentDetailController)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.ServerIP, cfg.ServerPort)))
}
