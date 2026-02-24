package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/fx"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg *config.Config, lc fx.Lifecycle) *Redis {
	logger.Info("Initializing Redis client...")
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Creds.Cache.Redis.Host, cfg.Creds.Cache.Redis.Port),
		Password:     cfg.Creds.Cache.Redis.Password,
		DB:           cfg.Creds.Cache.Redis.DB,
		PoolSize:     10,
		MinIdleConns: 2,
	})

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return rdb.Close()
		},
	})
	return &Redis{client: rdb}
}
