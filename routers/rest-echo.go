package main

import (
	"github.com/labstack/echo"
)

// Handler
func simple(c *echo.Context) error {
	c.String(200, "joł!")
	return nil
}

func main() {
	e := echo.New()

	e.Get("/item", simple)
	e.Get("/item/:id", simple)
	e.Get("/item/:id/subitem/:subId", simple)

	e.Run(":9090")
}
