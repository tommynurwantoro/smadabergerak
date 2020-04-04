package router

import (
	"tommynurwantoro/mygoweb/internal/domain"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, h domain.Handler) {
	e.GET("/", h.SmadaBergerak().Index)
	e.GET("/healthz", h.General().Healthz)
	e.GET("/transparansi", h.SmadaBergerak().Transparansi)
	e.GET("/progres", h.SmadaBergerak().Progres)
}
