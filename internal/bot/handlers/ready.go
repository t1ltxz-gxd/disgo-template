package handlers

import (
	"context"
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type readyHandler struct {
	client *bot.Client
	cfg    *config.Config
}

func NewReadyHandler(client *bot.Client, cfg *config.Config) types.Event {
	return &readyHandler{
		client: client,
		cfg:    cfg,
	}
}

func (h *readyHandler) Event() any {
	return &events.Ready{}
}

func (h *readyHandler) Handle(event any) error {
	ready := event.(*events.Ready)

	logger.Info("logged in",
		zap.String("username", ready.User.Username),
	)

	if !h.client.HasShardManager() {
		logger.Warn("shard manager is disabled")
		return nil
	}

	go func() {
		time.Sleep(2 * time.Second)
		setPresenceWithRetry(h.client, h.cfg, 3)
	}()

	return nil
}

func setPresenceWithRetry(client *bot.Client, cfg *config.Config, maxRetries int) {
	shardIDs := cfg.Bot.Sharding.ShardIDs
	if len(shardIDs) == 0 && cfg.Bot.Sharding.ShardCount > 0 {
		shardIDs = make([]int, cfg.Bot.Sharding.ShardCount)
		for i := 0; i < cfg.Bot.Sharding.ShardCount; i++ {
			shardIDs[i] = i
		}
	}

	for _, shardID := range shardIDs {
		var err error
		for attempt := 0; attempt < maxRetries; attempt++ {
			err = client.SetPresenceForShard(
				context.Background(),
				shardID,
				parseActivity(cfg),
				gateway.WithOnlineStatus(parseStatus(cfg.Bot.Status)),
			)
			if err == nil {
				break
			}
			if attempt < maxRetries-1 {
				time.Sleep(time.Duration(attempt+1) * time.Second)
			}
		}
		if err != nil {
			logger.Warn("Failed to set presence for shard after retries",
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

func parseActivity(cfg *config.Config) gateway.PresenceOpt {
	switch cfg.Bot.Activity.Type {
	case "playing":
		return gateway.WithPlayingActivity(cfg.Bot.Activity.Name)
	case "streaming":
		return gateway.WithStreamingActivity(cfg.Bot.Activity.Name, cfg.Bot.Activity.URL)
	case "listening":
		return gateway.WithListeningActivity(cfg.Bot.Activity.Name)
	case "watching":
		return gateway.WithWatchingActivity(cfg.Bot.Activity.Name)
	default:
		logger.Warn("Unknown activity type provided, defaulting to 'playing'", zap.String("activity_type", cfg.Bot.Activity.Type))
		return gateway.WithPlayingActivity(cfg.Bot.Activity.Name)
	}
}
