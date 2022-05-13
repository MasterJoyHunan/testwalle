package main

import (
	_ "embed"
	stderr "errors"
	"log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//go:embed Dockerfile
var dockerFile []byte

func main() {
	route := gin.Default()
	pprof.Register(route)
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

	route.GET("/error", func(c *gin.Context) {
		log.Printf("%+v", errors.Wrap(stderr.New("some err"), "local error"))
		c.JSON(200, gin.H{
			"code":    1,
			"message": string(dockerFile),
		})
	})

	log.Panic(route.Run("0.0.0.0:8820"))
}
