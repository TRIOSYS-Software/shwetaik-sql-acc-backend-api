package repositories

import (
	"shwetaik-sql-acc-backend-api/models"

	"gorm.io/gorm"
)

type PaymentDetailRepo struct {
	db *gorm.DB
}

func NewPaymentDetailRepo(db *gorm.DB) *PaymentDetailRepo {
	return &PaymentDetailRepo{db: db}
}

func (p *PaymentDetailRepo) GetAll() ([]models.PaymentDetail, error) {
	var paymentDetails []models.PaymentDetail
	err := p.db.Find(&paymentDetails).Error
	return paymentDetails, err
}
