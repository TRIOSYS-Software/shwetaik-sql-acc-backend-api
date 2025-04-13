package services

import (
	"shwetaik-sql-acc-backend-api/models"
	"shwetaik-sql-acc-backend-api/repositories"
)

type GLAccService struct {
	repo *repositories.GLAccRepo
}

func NewGLAccService(repo *repositories.GLAccRepo) *GLAccService {
	return &GLAccService{repo: repo}
}

func (p *GLAccService) GetAll() ([]models.GLAcc, error) {
	return p.repo.GetAll()
}

func (p *GLAccService) GetAllLowLevel() ([]models.GLAcc, error) {
	return p.repo.GetAllLowLevel()
}

func (p *GLAccService) FilterByCodes(code string) (*[]models.GLAcc, error) {
	return p.repo.FilterByCodes(code)
}
