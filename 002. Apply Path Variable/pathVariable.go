package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Student struct {
	id int
	name string
	age int
}

var studentTable = map[int]Student{
	1: {id: 1, name: "이진혁", age: 19},
	2: {id: 2, name: "이재혁", age: 18},
	3: {id: 3, name: "이준혁", age: 15},
}

func main() {
	router := gin.Default()

	router.GET("/students/:studentId", func(context *gin.Context) {
		studentId, _ := strconv.Atoi(context.Param("studentId"))
		student := studentTable[studentId]
		context.JSON(http.StatusOK, gin.H{
			"id": student.id,
			"name": student.name,
			"age": student.age,
		})
	})

	router.GET("/teachers", func(context *gin.Context) {
		/* anything */
	})

	router.GET("/teachers/*something", func(context *gin.Context) {
		somethingPath := context.Param("something")
		context.String(http.StatusOK, somethingPath)
	})

	router.Run()
}
