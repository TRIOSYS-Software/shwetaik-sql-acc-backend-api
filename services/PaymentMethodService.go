package services

import (
	"shwetaik-sql-acc-backend-api/dtos"
	"shwetaik-sql-acc-backend-api/models"
	"shwetaik-sql-acc-backend-api/repositories"
)

type PaymentMethodService struct {
	PaymentMethodRepo *repositories.PaymentMethodRepo
}

func NewPaymentMethodService(paymentMethodRepo *repositories.PaymentMethodRepo) *PaymentMethodService {
	return &PaymentMethodService{PaymentMethodRepo: paymentMethodRepo}
}

func (s *PaymentMethodService) GetAll() ([]dtos.PaymentMethodRequestDTO, error) {
	return s.PaymentMethodRepo.GetAll()
}

func (s *PaymentMethodService) GetByCode(code string) (*models.PaymentMethod, error) {
	return s.PaymentMethodRepo.GetByCode(code)
}
