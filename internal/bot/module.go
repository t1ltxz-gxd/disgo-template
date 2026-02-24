package bot

import (
	"go.t1ltxz.ninja/disgo-template/internal/bot/commands"
	"go.t1ltxz.ninja/disgo-template/internal/bot/components"
	"go.t1ltxz.ninja/disgo-template/internal/bot/handlers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"bot",
	fx.Provide(NewBot, NewHandler),
	commands.Module,
	handlers.Module,
	components.Module,
)
