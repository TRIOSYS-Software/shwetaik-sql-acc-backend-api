package controllers

import (
	"net/http"
	"shwetaik-sql-acc-backend-api/models"
	"shwetaik-sql-acc-backend-api/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentService *services.PaymentService
}

func NewPaymentController(paymentService *services.PaymentService) *PaymentController {
	return &PaymentController{PaymentService: paymentService}
}

func (p *PaymentController) GetAll(c echo.Context) error {
	payments, err := p.PaymentService.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, payments)
}

func (p *PaymentController) GetByDOCKEY(c echo.Context) error {
	docKey := c.Param("docKey")
	docKeyInt, err := strconv.Atoi(docKey)
	if err != nil {
		return err
	}
	payment, err := p.PaymentService.GetByDOCKEY(uint(docKeyInt))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Payment not found")
	}
	return c.JSON(http.StatusOK, payment)
}

func (p *PaymentController) Create(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return err
	}
	err := p.PaymentService.Create(payment)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, payment)
}
