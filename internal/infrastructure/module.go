package infrastructure

import (
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/cache"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/database"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure",
	cache.Module,
	database.Module,
	logger.Module,
)
