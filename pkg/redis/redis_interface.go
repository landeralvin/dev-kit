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
	GetCluster(ctx context.Context, key string) (string, error)
	SetCluster(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	DelCluster(ctx context.Context, keys ...string) error
	ExistsCluster(ctx context.Context, keys ...string) (int64, error)
	PingCluster(ctx context.Context) (string, error)
	CloseCluster(ctx context.Context) error
	HMSetCluster(ctx context.Context, key string, values map[string]interface{}) error
	HMGetCluster(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HGetAllCluster(ctx context.Context, key string) (map[string]string, error)
	HDelCluster(ctx context.Context, key string, fields ...string) error
	HGetCluster(ctx context.Context, key string, field string) (string, error)
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}

func (r *RedisClient) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.Client.Exists(ctx, keys...).Result()
}

func (r *RedisClient) Ping(ctx context.Context) (string, error) {
	return r.Client.Ping(ctx).Result()
}

func (r *RedisClient) Close(ctx context.Context) error {
	return r.Client.Close()
}

func (r *RedisClient) HMSet(ctx context.Context, key string, values map[string]interface{}) error {
	return r.Client.HMSet(ctx, key, values).Err()
}

func (r *RedisClient) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return r.Client.HMGet(ctx, key, fields...).Result()
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.Client.HGetAll(ctx, key).Result()
}

func (r *RedisClient) HDel(ctx context.Context, key string, fields ...string) error {
	return r.Client.HDel(ctx, key, fields...).Err()
}

func (r *RedisClient) HGet(ctx context.Context, key string, field string) (string, error) {
	return r.Client.HGet(ctx, key, field).Result()
}

func (r *RedisClient) GetCluster(ctx context.Context, key string) (string, error) {
	return r.ClientCluster.Get(ctx, key).Result()
}

func (r *RedisClient) SetCluster(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.ClientCluster.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) DelCluster(ctx context.Context, keys ...string) error {
	return r.ClientCluster.Del(ctx, keys...).Err()
}

func (r *RedisClient) ExistsCluster(ctx context.Context, keys ...string) (int64, error) {
	return r.ClientCluster.Exists(ctx, keys...).Result()
}

func (r *RedisClient) PingCluster(ctx context.Context) (string, error) {
	return r.ClientCluster.Ping(ctx).Result()
}

func (r *RedisClient) CloseCluster(ctx context.Context) error {
	return r.ClientCluster.Close()
}

func (r *RedisClient) HMSetCluster(ctx context.Context, key string, values map[string]interface{}) error {
	return r.ClientCluster.HMSet(ctx, key, values).Err()
}

func (r *RedisClient) HMGetCluster(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return r.ClientCluster.HMGet(ctx, key, fields...).Result()
}

func (r *RedisClient) HGetAllCluster(ctx context.Context, key string) (map[string]string, error) {
	return r.ClientCluster.HGetAll(ctx, key).Result()
}

func (r *RedisClient) HDelCluster(ctx context.Context, key string, fields ...string) error {
	return r.ClientCluster.HDel(ctx, key, fields...).Err()
}

func (r *RedisClient) HGetCluster(ctx context.Context, key string, field string) (string, error) {
	return r.ClientCluster.HGet(ctx, key, field).Result()
}
