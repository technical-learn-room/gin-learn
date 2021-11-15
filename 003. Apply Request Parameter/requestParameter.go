package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	id int
	name string
	age int
}

var database = map[int]student{
	1: {id: 1, name: "이진혁", age: 19},
	2: {id: 1, name: "이진혁", age: 18},
	3: {id: 2, name: "이재혁", age: 18},
	4: {id: 3, name: "이준혁", age: 15},
}

func main() {
	router := gin.Default()

	router.GET("/students", func(context *gin.Context) {
		name := context.Query("name")

		var response []student

		for _, s := range database {
			if s.name == name {
				response = append(response, s)
			}
		}

		context.JSON(http.StatusOK, response)
	})

	router.Run()
}