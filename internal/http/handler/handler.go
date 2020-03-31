package handler

import (
	"html/template"
	"tommynurwantoro/mygoweb/internal/domain"

	"github.com/gomarkdown/markdown"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	GeneralHandler       domain.GeneralHandler
	SmadaBergerakHandler domain.SmadaBergerakHandler
}

func (h *Handler) General() domain.GeneralHandler {
	return h.GeneralHandler
}

func (h *Handler) SmadaBergerak() domain.SmadaBergerakHandler {
	return h.SmadaBergerakHandler
}

func NewHandler(e *echo.Echo) domain.Handler {
	return &Handler{
		GeneralHandler:       NewGeneralHandler(),
		SmadaBergerakHandler: NewSmadaBergerakHandler(),
	}
}

// PRIVATE
func markDowner(args []byte) template.HTML {
	output := template.HTML(markdown.ToHTML(args, nil, nil))
	return template.HTML(output)
}
