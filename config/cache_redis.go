package config

import (
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

func ConnectRedis() *redis.Client {
	redisDefaultDb, err := strconv.Atoi(os.Getenv("REDIS_DEFAULT_DB"))

	if err != nil {
		panic("Error convert env REDIS_DEFAULT_DB")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       redisDefaultDb,              // use default DB
	})

	return rdb
}
