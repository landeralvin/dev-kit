package redis

import (
	"context"
	"time"
)

type RedisInterface interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (int64, error)
	Ping(ctx context.Context) (string, error)
	Close(ctx context.Context) error
	HMSet(ctx context.Context, key string, values map[string]interface{}) error
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) error
	HGet(ctx context.Context, key string, field string) (string, error)
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *RedisClient) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Exists(ctx, keys...).Result()
}

func (r *RedisClient) Ping(ctx context.Context) (string, error) {
	return r.client.Ping(ctx).Result()
}

func (r *RedisClient) Close(ctx context.Context) error {
	return r.client.Close()
}

func (r *RedisClient) HMSet(ctx context.Context, key string, values map[string]interface{}) error {
	return r.client.HMSet(ctx, key, values).Err()
}

func (r *RedisClient) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return r.client.HMGet(ctx, key, fields...).Result()
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

func (r *RedisClient) HDel(ctx context.Context, key string, fields ...string) error {
	return r.client.HDel(ctx, key, fields...).Err()
}

func (r *RedisClient) HGet(ctx context.Context, key string, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}
