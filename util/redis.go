package util

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func getRedisClient() *redis.Client {

	if redisClient != nil {
		return redisClient
	}
	settings := GetSettings()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:" + settings.RedisPort,
		Password: "",
		DB:       0,
	})
	return redisClient
}

type RehashCounter struct {
	client    *redis.Client
	keySuffix string
}

func (cli *RehashCounter) Init() {
	cli.client = getRedisClient()
	cli.keySuffix = "_Counter"
}

func (cli *RehashCounter) Add(key string, val int) error {
	key += cli.keySuffix
	return cli.client.Set(ctx, key, val, 0).Err()
}

func (cli *RehashCounter) Get(key string) (int, error) {
	key += cli.keySuffix
	val2, err := cli.client.Get(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	count, err := strconv.Atoi(val2)
	return count, err
}

func (cli *RehashCounter) Exist(key string) bool {
	exists, err := cli.client.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	if exists == 1 {
		return true
	}
	return false
}

type UrlCache struct {
	client    *redis.Client
	keySuffix string
}

func (storage *UrlCache) Init() {
	storage.client = getRedisClient()
	storage.keySuffix = "_UrlCache"
}

func (storage *UrlCache) Add(key, val string) error {
	key += storage.keySuffix
	return storage.client.Set(ctx, key, val, 0).Err()
}

func (storage *UrlCache) Get(key string) (string, error) {
	key += storage.keySuffix
	val2, err := storage.client.Get(ctx, key).Result()
	return val2, err
}

func (storage *UrlCache) Exist(key string) bool {
	exists, err := storage.client.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	if exists == 1 {
		return true
	}
	return false
}
