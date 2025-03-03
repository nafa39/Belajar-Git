package main

import (
	"rmt7-3rdparty/config"
	"rmt7-3rdparty/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/cities", func(c echo.Context) error {
		return handler.GetCities(c, cfg)
	})
	e.POST("/cek-ongkir", func(c echo.Context) error {
		return handler.GetOngkir(c, cfg)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
