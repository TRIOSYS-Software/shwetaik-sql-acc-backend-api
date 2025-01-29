package services

import (
	"shwetaik-sql-acc-backend-api/models"
	"shwetaik-sql-acc-backend-api/repositories"
)

type PaymentDetailService struct {
	PaymentDetailRepo *repositories.PaymentDetailRepo
}

func NewPaymentDetailService(paymentDetailRepo *repositories.PaymentDetailRepo) *PaymentDetailService {
	return &PaymentDetailService{PaymentDetailRepo: paymentDetailRepo}
}

func (s *PaymentDetailService) GetAll() ([]models.PaymentDetail, error) {
	return s.PaymentDetailRepo.GetAll()
}
