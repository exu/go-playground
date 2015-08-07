package main

import (
	"github.com/gin-gonic/gin"
)

// Handler
func simple(c *gin.Context) {
	c.String(200, "joł!")
}

func main() {
	r := gin.New()

	r.GET("/item", simple)
	r.GET("/item/:id", simple)
	r.GET("/item/:id/subitem/:subId", simple)

	r.Run(":8081")
}
