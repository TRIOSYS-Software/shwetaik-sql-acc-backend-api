package services

import (
	"shwetaik-sql-acc-backend-api/models"
	"shwetaik-sql-acc-backend-api/repositories"
)

type ProjectService struct {
	ProjectRepo *repositories.ProjectRepo
}

func NewProjectService(projectRepo *repositories.ProjectRepo) *ProjectService {
	return &ProjectService{ProjectRepo: projectRepo}
}

func (p *ProjectService) GetAll() ([]models.Project, error) {
	return p.ProjectRepo.GetAll()
}

func (p *ProjectService) GetByCode(code string) (*models.Project, error) {
	return p.ProjectRepo.GetByCode(code)
}
