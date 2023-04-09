package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "af4f2cd", "http://google.com", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "af4f2cd").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("af4f2cd", val)
}
