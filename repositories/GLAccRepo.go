package repositories

import (
	"shwetaik-sql-acc-backend-api/models"
	"strings"

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
	err := p.db.Where("NOT (PARENT = -1) AND CODE NOT LIKE '300-%' AND CODE NOT LIKE '400-%'").Find(&glAccs).Error
	return glAccs, err
}

func (p *GLAccRepo) GetByCode(code string) (*models.GLAcc, error) {
	var glAcc models.GLAcc
	err := p.db.Where("CODE = ?", code).First(&glAcc).Error
	return &glAcc, err
}

func (p *GLAccRepo) FilterByCodes(code string) (*[]models.GLAcc, error) {
	filter := strings.Split(code, ",")
	var glAccs []models.GLAcc
	if err := p.db.Where("CODE IN ?", filter).Find(&glAccs).Error; err != nil {
		return nil, err
	}
	dockey := []int{}
	for _, v := range glAccs {
		dockey = append(dockey, v.DOCKEY)
	}
	var FilteredGLAcc []models.GLAcc
	if err := p.db.Where("NOT PARENT = -1 AND NOT PARENT IN ? AND NOT DOCKEY IN ?", dockey, dockey).Find(&FilteredGLAcc).Error; err != nil {
		return nil, err
	}
	return &FilteredGLAcc, nil
}
