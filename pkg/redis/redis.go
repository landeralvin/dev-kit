package redis

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client        *redis.Client
	ClientCluster *redis.ClusterClient
}

type RedisClientConfig struct {
	Addr        string
	AddrCluster []string
	Password    string
	DB          int
}

type Option func(*redis.Options)

type OptionCluster func(*redis.ClusterOptions)

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

func ConnectWithOptionCluster(cfg RedisClientConfig) OptionCluster {
	return func(o *redis.ClusterOptions) {
		if len(cfg.AddrCluster) == 0 {
			cfg.AddrCluster = []string{"localhost:6379"}
		}

		if cfg.Password == "" {
			cfg.Password = ""
		}

		if cfg.DB == 0 {
			cfg.DB = 0
		}
	}
}

func NewRedisClusterClient(opts ...OptionCluster) *RedisClient {
	clusterOptions := &redis.ClusterOptions{}
	for _, opt := range opts {
		opt(clusterOptions)
	}
	return &RedisClient{
		ClientCluster: redis.NewClusterClient(clusterOptions),
	}
}

func NewRedisClient(opts ...Option) *RedisClient {
	options := &redis.Options{}
	for _, opt := range opts {
		opt(options)
	}
	return &RedisClient{
		Client: redis.NewClient(options),
	}
}
