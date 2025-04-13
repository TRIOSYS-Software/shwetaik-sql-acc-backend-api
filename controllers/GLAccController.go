package controllers

import (
	"net/http"
	"shwetaik-sql-acc-backend-api/services"

	"github.com/labstack/echo/v4"
)

type GLAccController struct {
	GLAccService *services.GLAccService
}

func NewGLAccController(GLAccService *services.GLAccService) *GLAccController {
	return &GLAccController{GLAccService: GLAccService}
}

func (p *GLAccController) GetAll(c echo.Context) error {
	GLAcc, err := p.GLAccService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, GLAcc)
}

func (p *GLAccController) GetAllLowLevel(c echo.Context) error {
	GLAcc, err := p.GLAccService.GetAllLowLevel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, GLAcc)
}

func (p *GLAccController) FilterByCodes(c echo.Context) error {
	code := c.QueryParam("codes")
	GLAcc, err := p.GLAccService.FilterByCodes(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, GLAcc)
}
