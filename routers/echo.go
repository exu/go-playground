package main

import (
	"math/rand"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	// mw "github.com/labstack/echo/middleware"
)

// Handler
func hello(c *echo.Context) error {
	return c.JSON(http.StatusOK, &map[string]interface{}{
		"nmel_errors_count": rand.Int(),
		"users":             rand.Int(),
	})

	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	e := echo.New()
	e.Get("/", hello)
	e.Run(":1323")
}
