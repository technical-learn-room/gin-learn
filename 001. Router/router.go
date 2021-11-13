package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		/* 구현 */
	})
	router.POST("/", func(context *gin.Context) {
		/* 구현 */
	})
	router.PUT("/", func(context *gin.Context) {
		/* 구현 */
	})
	router.PATCH("/", func(context *gin.Context) {
		/* 구현 */
	})
	router.DELETE("/", func(context *gin.Context) {
		/* 구현 */
	})
	router.HEAD("/", func(context *gin.Context) {
		/* 구현 */
	})
	router.OPTIONS("/", func(context *gin.Context) {
		/* 구현 */
	})

	router.GET("/hello", func(context *gin.Context) {
		fmt.Println("Hello, first function")
		time.Sleep(3 * time.Second)
	}, func(context *gin.Context) {
		fmt.Println("Hello, second function")
	}, func(context *gin.Context) {
		fmt.Println("Hello, third function")
	})

	router.Run()
}
