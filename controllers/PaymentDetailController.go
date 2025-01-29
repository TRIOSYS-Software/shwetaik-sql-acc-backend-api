package controllers

import (
	"net/http"
	"shwetaik-sql-acc-backend-api/services"

	"github.com/labstack/echo/v4"
)

type PaymentDetailController struct {
	PaymentDetailService *services.PaymentDetailService
}

func NewPaymentDetailController(paymentDetailService *services.PaymentDetailService) *PaymentDetailController {
	return &PaymentDetailController{PaymentDetailService: paymentDetailService}
}

func (p *PaymentDetailController) GetAll(c echo.Context) error {
	paymentDetails, err := p.PaymentDetailService.GetAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Payment details not found")
	}
	return c.JSON(http.StatusOK, paymentDetails)
}
