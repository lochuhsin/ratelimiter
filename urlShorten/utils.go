package UrlShorten

import (
	"fmt"
	"hash/crc32"
)

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
	return fmt.Sprintf("%08x", hashValue)[:7]
}
