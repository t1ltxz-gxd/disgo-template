package logger

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

type ZapSlogHandler struct {
	Zap *zap.Logger
}

func (h *ZapSlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *ZapSlogHandler) Handle(_ context.Context, record slog.Record) error {
	msg := record.Message
	switch record.Level {
	case slog.LevelDebug:
		h.Zap.Debug(msg)
	case slog.LevelInfo:
		h.Zap.Info(msg)
	case slog.LevelWarn:
		h.Zap.Warn(msg)
	case slog.LevelError:
		h.Zap.Error(msg)
	default:
		h.Zap.Info(msg)
	}
	return nil
}

func (h *ZapSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *ZapSlogHandler) WithGroup(name string) slog.Handler {
	return h
}
