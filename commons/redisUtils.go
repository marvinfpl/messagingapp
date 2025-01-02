package commons

import (
	"github.com/redis/go-redis/v9"
	"os"
	"context"
)

func InitRedis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URI"),
		Password: "",
		DB: 0,
	})
	err := rd.Ping(context.Background()).Err()
	if err != nil {
		panic("cant connect to redis: " + err.Error())
	}
	return rd
}