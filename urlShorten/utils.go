package UrlShorten

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"sync"
)

var UrlMapperInstance *UrlMapper

// move these to redis and database ...
type UrlMapper struct {
	urlMap map[string]string
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

func (mapper *UrlMapper) Init() {
	mapper.urlMap = make(map[string]string)
}

func (mapper *UrlMapper) Set(url string) string {
	outputs := strconv.Itoa(mapper.hash(url))
	mapper.urlMap[outputs] = url
	return outputs
}

func (mapper *UrlMapper) Get(url string) (string, bool) {
	if val, ok := mapper.urlMap[url]; ok {
		return val, ok
	}
	return "", false
}

func (mapper *UrlMapper) hash(url string) int {
	return len(url)
}

// var crcTable

var crcTable *crc32.Table

func getTable() *crc32.Table {
	if crcTable == nil {
		crcTable = crc32.MakeTable(crc32.Castagnoli)
	}
	return crcTable
}

func UrlHasher(url string) string {
	crc32q := getTable()
	hashValue := crc32.Checksum([]byte(url), crc32q)
	return fmt.Sprintf("%08x", hashValue)
}
