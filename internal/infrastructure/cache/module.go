package cache

import (
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/cache/redis"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"cache",
	redis.Module,
)
