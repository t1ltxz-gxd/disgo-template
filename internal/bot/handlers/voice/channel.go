package voice

import (
	"github.com/disgoorg/disgo/events"
	"go.t1ltxz.ninja/disgo-template/internal/bot/types"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type ChannelMovementLogger struct{}

var _ types.Event = (*ChannelMovementLogger)(nil)

func (c *ChannelMovementLogger) Event() any {
	return &events.GuildVoiceStateUpdate{}
}

func (c *ChannelMovementLogger) Handle(event any) error {
	vs := event.(*events.GuildVoiceStateUpdate)

	before := &vs.OldVoiceState
	after := &vs.VoiceState

	// Check if channel ID changed
	beforeChannelID := before.ChannelID
	afterChannelID := after.ChannelID

	if beforeChannelID == afterChannelID {
		return nil
	}

	username := vs.Member.User.Username

	switch {
	case (beforeChannelID == nil) && (afterChannelID != nil):
		logger.Info("User joined voice channel",
			zap.String("username", username),
			zap.String("channel_id", afterChannelID.String()),
		)

	case (beforeChannelID != nil) && (afterChannelID == nil):
		logger.Info("User left voice channel",
			zap.String("username", username),
			zap.String("channel_id", beforeChannelID.String()),
		)

	case beforeChannelID != nil:
		logger.Info("User switched voice channel",
			zap.String("username", username),
			zap.String("from_channel_id", beforeChannelID.String()),
			zap.String("to_channel_id", afterChannelID.String()),
		)
	}

	return nil
}

func (c *ChannelMovementLogger) Name() string {
	return "ChannelMovementLogger"
}

type ChannelMuteDeafLogger struct{}

var _ types.Event = (*ChannelMuteDeafLogger)(nil)

func (c *ChannelMuteDeafLogger) Event() any {
	return &events.GuildVoiceStateUpdate{}
}

func (c *ChannelMuteDeafLogger) Handle(event any) error {
	vs := event.(*events.GuildVoiceStateUpdate)

	before := &vs.OldVoiceState
	after := &vs.VoiceState

	afterChannelID := after.ChannelID
	if afterChannelID == nil {
		return nil
	}

	username := vs.Member.User.Username

	// Microphone mute state
	if before.SelfMute != after.SelfMute {
		if after.SelfMute {
			logger.Info("User muted microphone",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		} else {
			logger.Info("User unmuted microphone",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		}
	}

	// Server mute
	if before.GuildMute != after.GuildMute {
		if after.GuildMute {
			logger.Info("User was server-muted",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		} else {
			logger.Info("User was server-unmuted",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		}
	}

	// Self deaf
	if before.SelfDeaf != after.SelfDeaf {
		if after.SelfDeaf {
			logger.Info("User deafened themselves",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		} else {
			logger.Info("User undeafened themselves",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		}
	}

	// Server deaf
	if before.GuildDeaf != after.GuildDeaf {
		if after.GuildDeaf {
			logger.Info("User was server-deafened",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		} else {
			logger.Info("User was server-undeafened",
				zap.String("username", username),
				zap.String("channel_id", afterChannelID.String()),
			)
		}
	}

	return nil
}

func (c *ChannelMuteDeafLogger) Name() string {
	return "ChannelMuteDeafLogger"
}
