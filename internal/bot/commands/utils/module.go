package utils

import (
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"utils_commands",
	fx.Provide(
		fx.Annotate(
			NewPingCommand,
			fx.As(new(types.Command)),
			fx.ResultTags(`group:"commands"`),
		),
	),
)
