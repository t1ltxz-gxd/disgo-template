package bot

import (
	"context"
	"log/slog"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/sharding"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Bot struct {
	Client *bot.Client
	Cfg    *config.Config
}

func NewHandler() *handler.Mux {
	return handler.New()
}

type Params struct {
	fx.In

	Cfg     *config.Config
	Handler *handler.Mux
	Logger  *zap.Logger
	// Db      *database.Postres
}

type Result struct {
	fx.Out

	Bot    *Bot
	Client *bot.Client
}

func NewBot(p Params) (Result, error) {
	token := p.Cfg.Creds.Bot.Token

	client, err := disgo.New(
		token,
		bot.WithShardManagerConfigOpts(
			sharding.WithShardIDs(p.Cfg.Bot.Sharding.ShardIDs...),
			sharding.WithShardCount(p.Cfg.Bot.Sharding.ShardCount),
			sharding.WithAutoScaling(p.Cfg.Bot.Sharding.AutoScaling),
			sharding.WithGatewayConfigOpts(
				gateway.WithIntents(
					gateway.IntentGuilds,
					gateway.IntentGuildMessages,
					gateway.IntentDirectMessages,
					gateway.IntentGuildVoiceStates,
				),
			),
		),
		bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagGuilds)),
		bot.WithDefaultShardManager(),
		bot.WithLogger(slog.New(&logger.ZapSlogHandler{Zap: p.Logger})),
		bot.WithEventListeners(p.Handler),
	)
	if err != nil {
		logger.Fatal("Failed to create Discord session: " + err.Error())
	}

	b := &Bot{
		Client: client,
		Cfg:    p.Cfg,
	}

	return Result{
		Bot:    b,
		Client: client,
	}, nil
}

func (b *Bot) Serve(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting Discord bot...")

			err := b.Client.OpenShardManager(ctx)
			if err != nil {
				logger.Error("Failed to open Discord session", zap.Error(err))
				return err
			}

			logger.Info("Discord bot started successfully!")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping Discord bot...")

			b.Client.Close(ctx)

			logger.Info("Discord bot stopped")
			return nil
		},
	})
}
