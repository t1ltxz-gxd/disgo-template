package voice

import (
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"voice",
	fx.Provide(
		fx.Annotate(
			func() types.Event { return &ChannelMovementLogger{} },
			fx.ResultTags(`group:"handlers"`),
		),
		fx.Annotate(
			func() types.Event { return &ChannelMuteDeafLogger{} },
			fx.ResultTags(`group:"handlers"`),
		),
	),
)
