package services

import (
	"go.t1ltxz.ninja/disgo-template/internal/services/scheduler"
	"go.t1ltxz.ninja/disgo-template/internal/services/workerpool"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	scheduler.Module,
	workerpool.Module,
	fx.Invoke(initializeService),
)
