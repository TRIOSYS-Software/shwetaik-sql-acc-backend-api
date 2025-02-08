package repositories

import (
	"shwetaik-sql-acc-backend-api/models"

	"gorm.io/gorm"
)

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (p *ProjectRepo) GetAll() ([]models.Project, error) {
	var projects []models.Project
	err := p.db.Find(&projects).Error
	return projects, err
}

func (p *ProjectRepo) GetByCode(code string) (*models.Project, error) {
	var project models.Project
	err := p.db.Where("CODE = ?", code).First(&project).Error
	return &project, err
}
