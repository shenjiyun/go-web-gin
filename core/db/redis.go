package db

import "github.com/go-redis/redis/v8"

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6397",
		Password: "",
		DB:       0,
	})

}
