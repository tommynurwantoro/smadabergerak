package handler

import (
	"fmt"
	"net/http"
	"tommynurwantoro/mygoweb/internal/domain"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type SmadaBergerakHandler struct {
	Data domain.SmadaBergerakData
}

func (h *SmadaBergerakHandler) Index(c echo.Context) error {
	cc := c.(*domain.CustomContext)

	h.Data.TargetDonationCurrency = h.formatCurrency(cc.Config.TargetDonation)
	h.Data.TotalDonationCurrency = h.formatCurrency(cc.Config.TotalDonation)
	percentage := float64(cc.Config.TotalDonation) / float64(cc.Config.TargetDonation) * 100
	h.Data.Percentage = fmt.Sprintf("%.0f", percentage)

	return c.Render(http.StatusOK, "smada_bergerak", h.Data)
}

func NewSmadaBergerakHandler() domain.SmadaBergerakHandler {
	return &SmadaBergerakHandler{
		Data: domain.SmadaBergerakData{
			PageTitle: "SMADA Bergerak",
		},
	}
}

// PRIVATE
func (v *SmadaBergerakHandler) formatCurrency(amount int) string {
	m := message.NewPrinter(language.Indonesian)
	return m.Sprintf("Rp %d", amount)
}
