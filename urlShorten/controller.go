package UrlShorten

import (
	"net/http"
	"ratelimiter/util"
	
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = util.GetLogger("ratelimiter/controller")

type UrlShortentherHandler struct {
	rehashCounter util.RehashCounter
	urlCache util.UrlCache
}

func ApiEndpoints(r *gin.Engine) {

	counter, urlcache := util.RehashCounter{}, util.UrlCache{}
	counter.Init()
	urlcache.Init()
	h := &UrlShortentherHandler{
		rehashCounter: counter,
		urlCache: urlcache,
	}

	routes := r.Group("/urlShorten")
	routes.GET("/", h.GetOriginal)
	routes.POST("/", h.CreateUrl)

}

func (handler *UrlShortentherHandler) CreateUrl(c *gin.Context) {

	url := Url{}
	if c.BindJSON(&url) == nil {
		logger.Info(c)
	}
	logger.Info(url)

	rawHashUrl := UrlHasher(url.OriginalUrl)
	// get first seven, as seven is a magic number, will discuss specifically
	rawHashUrl = rawHashUrl[:7]
	// 1. check if rawHashUrl is in redis && database else store in redis and db
	rehashCount := 0

	if !handler.urlCache.Exist(rawHashUrl){
		handler.urlCache.Add(rawHashUrl, url.OriginalUrl)
		handler.rehashCounter.Add(rawHashUrl, rehashCount)
		url.ShortenUrl = rawHashUrl
	}


	// 2. if so add predefined string and rehash once, add counter by one to restore original url

	// 3. store everything in redis (expire time) and postgres

	// 4. Add default rate limit policy for this url (maybe in elasticsearch)

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
