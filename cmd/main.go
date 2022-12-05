package main

import (
	"go-be-service/internal/hello1"
	"go-be-service/internal/hello2"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello1", hello1.Hello1)
	router.GET("/hello2", hello2.Hello2)
	err := router.Run(":2411")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
