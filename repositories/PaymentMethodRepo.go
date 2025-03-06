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
	err := p.db.Raw("SELECT pm.*, gl.DESCRIPTION FROM PMMETHOD pm JOIN GL_ACC gl ON pm.CODE = gl.CODE").Scan(&paymentMethods).Error
	// err := p.db.Joins("JOIN GL_ACC gl ON gl.code = PMMETHOD.code").Find(&paymentMethods).Error
	return paymentMethods, err
}

func (p *PaymentMethodRepo) GetByCode(code string) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := p.db.Where("CODE = ?", code).First(&paymentMethod).Error
	return &paymentMethod, err
}
