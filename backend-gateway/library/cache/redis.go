package cache

import (
	"time"

	"backend-gateway/library/config"

	logger "github.com/cihub/seelog"
	"github.com/go-redis/redis/v7"
)

func NewRedis(c *config.RedisConfig) *redis.Client {
	poolTimeout, err := time.ParseDuration(c.PoolTimeout)
	if err != nil {
		poolTimeout = 30 * time.Second
	}

	idleTimeout, err := time.ParseDuration(c.IdleTimeout)
	if err != nil {
		idleTimeout = 1 * time.Minute
	}

	redis := redis.NewClient(&redis.Options{
		Addr:        c.Addresses[c.AddressIndex],
		Password:    c.Credential,
		DB:          c.Db,
		PoolSize:    c.PoolSize,
		PoolTimeout: poolTimeout,
		IdleTimeout: idleTimeout,
	})

	// 测试 Redis 服务连接情况
	status, err := redis.Ping().Result()
	if err != nil {
		logger.Errorf("connect to redis server error: %+v, status: %s", err, status)
		panic(err)
	}

	logger.Infof("connect to redis server: %s, status: %s", c.Addresses[c.AddressIndex], status)

	return redis
}
