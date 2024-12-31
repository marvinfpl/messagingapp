package repository

import (
	"github.com/redis/go-redis/v9"
	"messagingapp/commons"
	"context"
	"errors"
)

type RedisRepository struct {
	client *redis.Client
}

func (r *RedisRepository) NewRedisRepository() *RedisRepository {
	return &RedisRepository{
		client: commons.InitRedis(),
	}
}

func (r *RedisRepository) AddWebsocketConnection(userKey string, userValue string) error {
	err := r.client.Set(context.Background(), userKey, userValue, 0).Err()
	return err
}

func (r *RedisRepository) RemoveWebsocketConnection(userKey string) error {
	err := r.client.Del(context.Background(), userKey).Err()
	return err
}

func (r *RedisRepository) GetWebsocketConnection(userKey string) (string, error) {
	conn, err := r.client.Get(context.Background(), userKey).Result()
	if err == redis.Nil {
		return "", errors.New("key doesn't exists:" + err.Error())
	}
	return conn, err
}