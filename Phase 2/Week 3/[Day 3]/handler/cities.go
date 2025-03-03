package handler

import (
	"github.com/labstack/echo/v4"

	"3rd-party/config"
	"3rd-party/utils"
)

func GetCities(c echo.Context, cfg *config.Config) error {
	newUrl := cfg.RajaOngkirAPIURL
	headers := map[string]string{
		"key": cfg.RajaOngkirAPIKey,
	}

	resp, err := utils.RequestGet(newUrl, headers)
	return nil
}
