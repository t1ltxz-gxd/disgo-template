package commands

import (
	"go.t1ltxz.ninja/disgo-template/internal/bot/commands/utils"
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"commands",
	fx.Invoke(NewRegistry),
	fx.Provide(
		fx.Annotate(
			NewTestCommand,
			fx.As(new(types.Command)),
			fx.ResultTags(`group:"commands"`),
		),
	),
	utils.Module,
)
