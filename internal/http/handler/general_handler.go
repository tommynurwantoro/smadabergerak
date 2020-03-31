package handler

import (
	"net/http"
	"tommynurwantoro/mygoweb/internal/domain"

	"github.com/labstack/echo/v4"
)

type GeneralHandler struct{}

func (g *GeneralHandler) Healthz(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func NewGeneralHandler() domain.GeneralHandler {
	return &GeneralHandler{}
}
