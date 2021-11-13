package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	defaultRouter := gin.Default()
	defaultRouter.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})
	defaultRouter.Run()
}
