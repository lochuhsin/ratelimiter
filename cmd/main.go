package main

import (
	"net/http"
	"ratelimiter/UrlShorten"
	"ratelimiter/util"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log := util.GetLogger("main")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	UrlShorten.ApiEndpoints(r)
	settings := util.GetSettings()
	log.Info("Basic settings")
	log.Info(settings)
	r.Run()
}
