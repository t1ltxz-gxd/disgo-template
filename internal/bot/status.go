package bot

import (
	"context"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func (b *Bot) updateStatus(ctx context.Context) {
	if !b.Client.HasShardManager() {
		logger.Warn("shard manager is disabled")
		return
	}

	shardIDs := b.Cfg.Bot.Sharding.ShardIDs
	if len(shardIDs) == 0 && b.Cfg.Bot.Sharding.ShardCount > 0 {
		shardIDs = make([]int, b.Cfg.Bot.Sharding.ShardCount)
		for i := 0; i < b.Cfg.Bot.Sharding.ShardCount; i++ {
			shardIDs[i] = i
		}
	}

	for _, shardID := range shardIDs {
		err := b.Client.SetPresenceForShard(
			ctx,
			shardID,
			parseActivityType(b.Cfg),
			gateway.WithOnlineStatus(parseStatus(b.Cfg.Bot.Status)),
		)
		if err != nil {
			logger.Error("Failed to set presence for shard",
				zap.Int("shard_id", shardID),
				zap.Error(err),
			)
		}
	}
}

func parseStatus(status string) discord.OnlineStatus {
	switch status {
	case "online":
		return discord.OnlineStatusOnline
	case "idle":
		return discord.OnlineStatusIdle
	case "dnd":
		return discord.OnlineStatusDND
	case "invisible":
		return discord.OnlineStatusInvisible
	default:
		logger.Warn("Unknown status provided, defaulting to 'online'", zap.String("status", status))
		return discord.OnlineStatusOnline
	}
}

func parseActivityType(cfg *config.Config) gateway.PresenceOpt {
	switch cfg.Bot.Activity.Type {
	case "playing":
		return gateway.WithPlayingActivity(cfg.Bot.Activity.Name)
	case "streaming":
		return gateway.WithStreamingActivity(cfg.Bot.Activity.Name, cfg.Bot.Activity.URL)
	case "listening":
		return gateway.WithListeningActivity(cfg.Bot.Activity.Name)
	case "watching":
		return gateway.WithWatchingActivity(cfg.Bot.Activity.Name)
	case "competing":
		return gateway.WithCompetingActivity(cfg.Bot.Activity.Name)
	default:
		logger.Warn("Unknown activity type provided, defaulting to 'playing'", zap.String("activity_type", cfg.Bot.Activity.Name))
		return gateway.WithPlayingActivity(cfg.Bot.Activity.Name)
	}
}
