package main

import (
	"go-xendit-rmt003/entity"
	"go-xendit-rmt003/helper"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func OrderHandler(c echo.Context) error {

	orderRequest := new(entity.OrderRequest)

	// ini contoh code save data to database
	if err := c.Bind(orderRequest); err != nil {
		return err
	}

	invoiceRes, err := helper.CreateInvoice(orderRequest.Product, orderRequest.Customer)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error while creating invoices")
	}

	return c.JSON(http.StatusOK, invoiceRes)
}

func main() {
	e := echo.New()

	//router order handler
	e.POST("/order", OrderHandler)

	//router invoice handler
	e.POST("/invoice", helper.CreateInvoice())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Start(":" + port)
}
