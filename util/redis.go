package util

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func getRedisClient() *redis.Client {

	if redisClient != nil {
		return redisClient
	}
	redisPort := os.Getenv("REDIS_PORT")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:" + redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return redisClient
}

type RehashCounter struct {
	client *redis.Client
}

func (cli *RehashCounter) Init() {
	cli.client = getRedisClient()
}

func (cli *RehashCounter) Add(key string, val int) error {
	return cli.client.Set(ctx, key, val, 0).Err()
}

func (storage *RehashCounter) Get(key string) (string, error) {
	val2, err := storage.client.Get(ctx, "key2").Result()
	return val2, err
}

func (cli *RehashCounter) Exist(key string) bool { return false }

type UrlCache struct {
	client *redis.Client
}

func (storage *UrlCache) Init() {
	storage.client = getRedisClient()
}

func (storage *UrlCache) Add(key, val string) error {
	return storage.client.Set(ctx, key, val, 0).Err()
}

func (storage *UrlCache) Get(key string) (string, error) {
	val2, err := storage.client.Get(ctx, "key2").Result()
	return val2, err
}

func (storage *UrlCache) Exist(key string) bool { return false }
