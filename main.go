package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 1,
			"message": "pong",
		})
	})

	log.Panic(route.Run("0.0.0.0:8888"))
}

