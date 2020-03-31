package domain

import (
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Config *StaticConfig
}

type StaticConfig struct {
	TargetDonation int `yaml:"targetDonation"`
	TotalDonation  int `yaml:"totalDonation"`
}
