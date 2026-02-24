package scheduler

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"scheduler",
	fx.Provide(New),
	fx.Invoke(registerLifecycle),
)

func registerLifecycle(lc fx.Lifecycle, s *Scheduler) {
	lc.Append(fx.Hook{
		OnStart: s.Start,
		OnStop:  s.Stop,
	})
}
