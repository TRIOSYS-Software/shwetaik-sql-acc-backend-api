package controllers

import (
	"fmt"
	"net/http"
	"shwetaik-sql-acc-backend-api/services"

	"github.com/labstack/echo/v4"
)

type ProjectController struct {
	ProjectService *services.ProjectService
}

func NewProjectController(projectService *services.ProjectService) *ProjectController {
	return &ProjectController{ProjectService: projectService}
}

func (p *ProjectController) GetAll(c echo.Context) error {
	projects, err := p.ProjectService.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, projects)
}

func (p *ProjectController) GetByCode(c echo.Context) error {
	code := c.Param("code")
	project, err := p.ProjectService.GetByCode(code)
	if err != nil {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Project with code %s not found", code))
	}
	return c.JSON(http.StatusOK, project)
}
