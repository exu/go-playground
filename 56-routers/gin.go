package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, &map[string]interface{}{"aaa": 1})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
