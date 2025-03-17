package main

import (
	"fmt"
	"os"
	"strings"

	"shwetaik-sql-acc-backend-api/config"
	"shwetaik-sql-acc-backend-api/controllers"
	"shwetaik-sql-acc-backend-api/repositories"
	"shwetaik-sql-acc-backend-api/routes"
	"shwetaik-sql-acc-backend-api/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	f, _ := os.OpenFile("echo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	e := echo.New()
	cfg, err := config.GetConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, level=${level}, remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, error=${error}, bytes_in=${bytes_in}, bytes_out=${bytes_out}\n",
		Output: f,
	}))
	e.Use(middleware.Recover())
	origins := strings.Split(cfg.AllowedOrigin, ",")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "shwetaik"},
	}))

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

	projectRepo := repositories.NewProjectRepo(db)
	projectService := services.NewProjectService(projectRepo)
	projectController := controllers.NewProjectController(projectService)
	routes.ProjectRoutes(apiV1, projectController)

	glAccRepo := repositories.NewGLAccRepo(db)
	glAccService := services.NewGLAccService(glAccRepo)
	glAccController := controllers.NewGLAccController(glAccService)
	routes.GLAccRoutes(apiV1, glAccController)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.ServerIP, cfg.ServerPort)))
}
