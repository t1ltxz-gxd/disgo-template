package main

import (
	"go.t1ltxz.ninja/disgo-template/internal/bot"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure"
	"go.t1ltxz.ninja/disgo-template/internal/repository"
	"go.t1ltxz.ninja/disgo-template/internal/services"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		bot.Module,
		config.Module,
		infrastructure.Module,
		repository.Module,
		services.Module,

		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(func(lc fx.Lifecycle, b *bot.Bot) {
			b.Serve(lc)
		}),
	).Run()
}
