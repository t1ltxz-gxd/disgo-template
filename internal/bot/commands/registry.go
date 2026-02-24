package commands

import (
	"context"
	"strings"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"

	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Client    *bot.Client
	Commands  []types.Command `group:"commands"`
	Config    *config.Config
}

func NewRegistry(p Params) {
	r := handler.New()

	var creates []discord.ApplicationCommandCreate

	for _, cmd := range p.Commands {
		creates = append(creates, cmd.Command())

		r.SlashCommand("/"+cmd.Name(), cmd.Handle)
	}

	p.Client.EventManager.AddEventListeners(r)

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Registering commands...")

			if strings.ToLower(p.Config.Creds.Env) == "dev" && p.Config.Bot.TestGuildID != "" {
				guildID, err := snowflake.Parse(p.Config.Bot.TestGuildID)
				if err != nil {
					logger.Error("failed to parse guild id", zap.Error(err))
					return err
				}
				logger.Info("Syncing commands to test guild (development mode)", zap.String("guild_id", p.Config.Bot.TestGuildID))

				_, err = p.Client.Rest.SetGuildCommands(
					p.Client.ApplicationID,
					guildID,
					creates,
				)
				if err != nil {
					logger.Error("failed to register guild commands", zap.Error(err))
					return err
				}

				logger.Info("Clearing global commands to avoid duplication")
				_, errGlobal := p.Client.Rest.SetGlobalCommands(
					p.Client.ApplicationID,
					[]discord.ApplicationCommandCreate{},
				)
				if errGlobal != nil {
					logger.Warn("failed to clear global commands", zap.Error(errGlobal))
				}

				return nil
			}

			logger.Info("Syncing commands globally (production mode)")
			_, err := p.Client.Rest.SetGlobalCommands(
				p.Client.ApplicationID,
				creates,
			)
			if err != nil {
				logger.Error("failed to register global commands", zap.Error(err))
				return err
			}

			return nil
		},
	})
}
