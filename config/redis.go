package config

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/lv123123long/pine/global"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
		Password: "",
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis")
	}

	global.RedisDB = RedisClient
}