package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-blog/config"
	"log"
)

func Connect() *redis.Client {
	host := config.Get("redis.host")
	port := config.Get("redis.port")
	password := config.Get("redis.password")

	if host == "" || port == "" {
		log.Println("can not connect to redis, config error")
		return nil
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
}

func Close(client *redis.Client) {
	client.Close()
}
