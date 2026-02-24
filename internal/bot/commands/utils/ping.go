package utils

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.t1ltxz.ninja/disgo-template/internal/utils"
	"go.uber.org/zap"
)

type PingCommand struct {
	cfg *config.Config
}

func (p *PingCommand) DevOnly() bool {
	return false
}

func NewPingCommand(cfg *config.Config) types.Command {
	return &PingCommand{
		cfg: cfg,
	}
}

func (p *PingCommand) Name() string {
	return "ping"
}

func (p *PingCommand) Command() discord.ApplicationCommandCreate {
	return discord.SlashCommandCreate{
		Name:        p.Name(),
		Description: "Check the bot's latency",
	}
}

func (p *PingCommand) Handle(
	_ discord.SlashCommandInteractionData,
	e *handler.CommandEvent,
) error {

	color, err := utils.HexToRGBInt(p.cfg.Bot.Color.Success)
	if err != nil {
		logger.Warn("Failed to parse color, using default", zap.Error(err))
	}

	embed := discord.NewEmbedBuilder().
		SetTitle("🏓 Pong!").
		SetColor(color).
		Build()

	if err := e.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{embed},
	}); err != nil {
		logger.Error("failed to send pong:", zap.Error(err))
		return err
	}

	return nil
}

var _ types.Command = (*PingCommand)(nil)
