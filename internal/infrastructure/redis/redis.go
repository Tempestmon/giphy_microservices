package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisKey struct {
	Key   string
	Value string
}

type RedisClient struct {
	address string
	client  *redis.Client
	context context.Context
}

type Redis interface {
	GetKey(key string) RedisKey
	SetKey(key string, value string) error
}

func NewRedisClient(address string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	return &RedisClient{address: address, client: rdb, context: context.Background()}
}

func (r *RedisClient) GetKey(key string) RedisKey {
	value, err := r.client.Get(r.context, key).Result()
	if err != nil {
		return RedisKey{}
	}
	return RedisKey{Key: key, Value: value}
}

func (r *RedisClient) SetKey(key string, value string) error {
	r.client.Set(r.context, key, value, 200)
	return nil
}
