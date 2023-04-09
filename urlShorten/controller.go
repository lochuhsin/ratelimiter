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
// TODO: add validator, store data to postgres, add rate limit default policy
func (handler *UrlShortentherHandler) CreateUrl(c *gin.Context) {
	settings := util.GetSettings()
	url := Url{}
	if c.BindJSON(&url) == nil {
		logger.Info(c)
	}
	originalUrl := url.OriginalUrl
	hashUrl := ""
	rehashCount := 0
	for {
		hashUrl = UrlHasher(originalUrl)
		if !handler.urlCache.Exist(hashUrl){
			break
		}
		originalUrl += settings.PredefineString
		rehashCount ++
	}

	err := handler.urlCache.Add(hashUrl, originalUrl)
	if err != nil{
		panic(err)
	}
	err = handler.rehashCounter.Add(hashUrl, rehashCount)
	if err != nil{
		panic(err)
	}
	url.ShortenUrl = hashUrl
	c.JSON(http.StatusOK, url)
}

// TODO: fix this c.JSON using struct to bind values
// fix url return with redundant "" -> """yahoo.com"
func (handler *UrlShortentherHandler) GetOriginal(c *gin.Context) {
	shortenUrl := c.DefaultQuery("shortenUrl", "")
	settings := util.GetSettings()
	if len(shortenUrl) == 0{
		c.JSON(http.StatusOK, "")
	}

	if !handler.urlCache.Exist(shortenUrl){
		c.JSON(http.StatusBadRequest, "")
	}

	hashedUrl, err := handler.urlCache.Get(shortenUrl)
	if err != nil{
		panic(err)
	}

	count, err := handler.rehashCounter.Get(shortenUrl)
	if err != nil{
		panic(err)
	}

	upperBound := len(hashedUrl) - len(settings.PredefineString)*count
	c.JSON(http.StatusOK, string(hashedUrl[:upperBound]))
}
