package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed Dockerfile
var dockerFile []byte

func main() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "pong",
		})
	})

	route.GET("/getFile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    1,
			"message": string(dockerFile),
		})
	})

	time.After(1 * time.Second)
	log.Panic(route.Run("0.0.0.0:8821"))
}
