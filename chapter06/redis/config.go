package redis

import (
	"os"

	"gopkg.in/redis.v5"
)

// Setup 함수는 redis 클라이언트를 초기화한다.
func Setup() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDISPASSWORD"),
		DB:       0, // user default DB
	})

	_, err := client.Ping().Result()

	return client, err
}
