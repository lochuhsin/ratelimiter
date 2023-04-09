package main

import (
	"net/http"
	"os"
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
	port := os.Getenv("PORT")

	log.Info(port)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
