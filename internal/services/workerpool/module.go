package workerpool

import "go.uber.org/fx"

var Module = fx.Module(
	"workerpool",
	fx.Provide(NewWorkerPool),
)
