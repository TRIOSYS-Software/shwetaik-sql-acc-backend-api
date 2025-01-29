package controllers

import (
	"net/http"
	"shwetaik-sql-acc-backend-api/services"

	"github.com/labstack/echo/v4"
)

type PaymentMethodController struct {
	PaymentMethodService *services.PaymentMethodService
}

func NewPaymentMethodController(paymentMethodService *services.PaymentMethodService) *PaymentMethodController {
	return &PaymentMethodController{PaymentMethodService: paymentMethodService}
}

func (p *PaymentMethodController) GetAll(c echo.Context) error {
	paymentMethods, err := p.PaymentMethodService.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, paymentMethods)
}

func (p *PaymentMethodController) GetByCode(c echo.Context) error {
	code := c.Param("code")
	paymentMethod, err := p.PaymentMethodService.GetByCode(code)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, paymentMethod)
}
