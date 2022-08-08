package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			c.JSON(
				http.StatusAccepted, gin.H{
					"s": "as",
				})
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "Status OK",
		})
	})
	route.Run()
}
