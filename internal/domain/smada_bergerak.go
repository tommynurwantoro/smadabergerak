package domain

import (
	"github.com/labstack/echo/v4"
)

type SmadaBergerakData struct {
	PageTitle              string
	TargetDonationCurrency string
	TotalDonationCurrency  string
	Percentage             string
}

type SmadaBergerakHandler interface {
	Index(c echo.Context) error
	Transparansi(c echo.Context) error
	Progres(c echo.Context) error
}
