package configs

import (
	"github.com/hoxito/mysql_comparer/tools/env"

	"github.com/go-redis/redis/v7"
)

//Client instance

func Client() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     env.Get().RedisURL,
		Password: "",
		DB:       0,
	})
}
