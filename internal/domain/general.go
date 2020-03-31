package domain

import "github.com/labstack/echo/v4"

type GeneralHandler interface {
	Healthz(c echo.Context) error
}
