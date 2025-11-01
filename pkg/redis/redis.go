package redis

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

type RedisClientConfig struct {
	Addr     string
	Password string
	DB       int
}

type Option func(*redis.Options)

func ConnectWithOption(cfg RedisClientConfig) Option {
	return func(o *redis.Options) {
		if cfg.Addr == "" {
			cfg.Addr = "localhost:6379"
		}

		if cfg.Password == "" {
			cfg.Password = ""
		}

		if cfg.DB == 0 {
			cfg.DB = 0
		}
	}
}

func NewRedisClient(opts ...Option) *RedisClient {
	options := &redis.Options{}
	for _, opt := range opts {
		opt(options)
	}
	return &RedisClient{
		client: redis.NewClient(options),
	}
}
