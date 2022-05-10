package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api/version", func (c echo.Context) error {
		return c.JSON(200,&echo.Map{"version": "1.0"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}