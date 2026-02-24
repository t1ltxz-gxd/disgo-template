package handlers

import (
	"go.t1ltxz.ninja/disgo-template/internal/bot/handlers/voice"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers",
	fx.Provide(
		fx.Annotate(
			NewReadyHandler,
			fx.ResultTags(`group:"handlers"`),
		),
	),
	fx.Invoke(NewRegistry),
	voice.Module,
)
