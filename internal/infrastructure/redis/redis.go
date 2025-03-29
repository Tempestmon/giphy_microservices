package redis

import (
	"context"
	"errors"
	"fmt"
	"giphy_microservices/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisKey struct {
	Key   string
	Value string
}

var ErrKeyNotFound = errors.New("Redis key not found")

type RedisClient struct {
	client  *redis.Client
	context context.Context
}

type Redis interface {
	GetKey(key string) RedisKey
	SetKey(key string, value string) error
}

func NewRedisClient(cfg *config.RedisConfig) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	return &RedisClient{
		client:  rdb,
		context: context.Background(),
	}, nil
}

func (r *RedisClient) GetKey(key string) (string, error) {
	value, err := r.client.Get(r.context, key).Result()
	fmt.Printf("redis get key error %s", err)
	if err != nil {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (r *RedisClient) SetKey(key string, value string) (string, error) {
	return r.client.Set(r.context, key, value, time.Hour).Result()
}
