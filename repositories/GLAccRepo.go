package repositories

import (
	"shwetaik-sql-acc-backend-api/models"

	"gorm.io/gorm"
)

type GLAccRepo struct {
	db *gorm.DB
}

func NewGLAccRepo(db *gorm.DB) *GLAccRepo {
	return &GLAccRepo{db: db}
}

func (p *GLAccRepo) GetAll() ([]models.GLAcc, error) {
	var glAccs []models.GLAcc
	err := p.db.Find(&glAccs).Error
	return glAccs, err
}

func (p *GLAccRepo) GetAllLowLevel() ([]models.GLAcc, error) {
	var glAccs []models.GLAcc
	err := p.db.Where("NOT (PARENT = -1) AND CODE NOT LIKE '4%%-%'").Find(&glAccs).Error
	return glAccs, err
}
