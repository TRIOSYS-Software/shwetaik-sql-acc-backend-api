package repositories

import (
	"shwetaik-sql-acc-backend-api/models"

	"gorm.io/gorm"
)

type PaymentMethodRepo struct {
	db *gorm.DB
}

func NewPaymentMethodRepo(db *gorm.DB) *PaymentMethodRepo {
	return &PaymentMethodRepo{db: db}
}

func (p *PaymentMethodRepo) GetAll() ([]models.PaymentMethod, error) {
	var paymentMethods []models.PaymentMethod
	err := p.db.Find(&paymentMethods).Error
	return paymentMethods, err
}

func (p *PaymentMethodRepo) GetByCode(code string) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := p.db.Where("CODE = ?", code).First(&paymentMethod).Error
	return &paymentMethod, err
}
