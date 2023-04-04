package ratelimit

import (
	"strconv"
	"sync"
)

var UrlMapperInstance *UrlMapper

// move these to redis and database ...
type UrlMapper struct {
	urlMap map[string]string
	mutex  sync.Mutex
}

var locker sync.Mutex

func GetMapperInstance() *UrlMapper {
	if UrlMapperInstance == nil {
		locker.Lock()
		if UrlMapperInstance == nil {
			defer locker.Unlock()
			mapper := UrlMapper{}
			mapper.Init()
			UrlMapperInstance = &mapper
		}
	}
	return UrlMapperInstance
}

func (self *UrlMapper) Init() {
	self.urlMap = make(map[string]string)
}

func (self *UrlMapper) Set(url string) string {
	outputs := strconv.Itoa(self.hash(url))
	self.urlMap[outputs] = url
	return outputs
}

func (self *UrlMapper) Get(url string) (string, bool) {
	if val, ok := self.urlMap[url]; ok {
		return val, ok
	}
	return "", false
}

func (self *UrlMapper) hash(url string) int {
	return len(url)
}
