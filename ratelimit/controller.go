package ratelimit

import (
	"net/http"
	"ratelimiter/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UrlShortentherHandler struct {
}

var logger *logrus.Logger = util.GetLogger("ratelimiter/controller")

func ApiEndpoints(r *gin.Engine) {
	h := &UrlShortentherHandler{}

	routes := r.Group("/urlShorten")
	routes.GET("/", h.GetOriginal)
	routes.POST("/", h.CreateUrl)

}

func (handler *UrlShortentherHandler) CreateUrl(c *gin.Context) {
	mapper := GetMapperInstance()

	url := Url{}
	if c.BindJSON(&url) == nil {
		logger.Info(c)
	}
	logger.Info(url)

	url.ShortenUrl = mapper.Set(url.OriginalUrl)

	c.JSON(http.StatusOK, url)
}

func (handler *UrlShortentherHandler) GetOriginal(c *gin.Context) {
	shortenUrl := c.DefaultQuery("shortenUrl", "")
	mapper := GetMapperInstance()
	url := Url{}

	if val, status := mapper.Get(shortenUrl); status {
		url.OriginalUrl = val
	}
	url.ShortenUrl = shortenUrl
	logger.Info(url)
	c.JSON(http.StatusOK, url)
}
