package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomErrorMessages struct {
	userMessage     string
	internalMessage string
}

func main() {
	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := "Internal Server Error"
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Message.(string)
			log.Println(he.Internal)
		}

		c.JSON(code, map[string]string{
			"error": msg,
		})
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/error", func(c echo.Context) error {
		// errorMessage := CustomErrorMessages{
		// 	userMessage:     "Bad Request",
		// 	internalMessage: "[main][controller][/error] Bad Request",
		// }
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong username or password").WithInternal(errors.New("Bad Request"))
	})

	e.GET("/panic", func(c echo.Context) error {
		panic("panic")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
