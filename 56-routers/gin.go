package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, &map[string]interface{}{
			"nmel_errors_count": rand.Int(),
			"users":             rand.Int(),
		})
	})

	r.Run(":8081")
}
