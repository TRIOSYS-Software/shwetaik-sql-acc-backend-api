package services

import (
	"shwetaik-sql-acc-backend-api/models"
	"shwetaik-sql-acc-backend-api/repositories"
)

type PaymentService struct {
	PaymentRepo *repositories.PaymentRepo
}

func NewPaymentService(paymentRepo *repositories.PaymentRepo) *PaymentService {
	return &PaymentService{PaymentRepo: paymentRepo}
}

func (s *PaymentService) GetAll() ([]models.Payment, error) {
	return s.PaymentRepo.GetAll()
}

func (s *PaymentService) GetByDOCKEY(docKey uint) (*models.Payment, error) {
	return s.PaymentRepo.GetByDOCKEY(docKey)
}

func (s *PaymentService) Create(payment *models.Payment) error {
	return s.PaymentRepo.Create(payment)
}
