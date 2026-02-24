package logger

import (
	"go.t1ltxz.ninja/disgo-template/internal/config"
	"go.t1ltxz.ninja/disgo-template/internal/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Module provides dependencies for working with logging
var Module = fx.Module(
	"logger",
	fx.Provide(NewLogger),
)

// NewLogger creates a new logger instance based on the configuration
func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	env := cfg.Creds.Env
	builder := NewZapLoggerBuilder().WithEnv(env).WithName(cfg.Logger.Name)
	if env == "production" || env == "prd" || env == "prod" {
		size, err := utils.ParseBytes(cfg.Logger.FileSyncer.MaxSize)
		if err != nil {
			return nil, err
		}
		fileConf := &lumberjack.Logger{
			Filename:   cfg.Logger.FileSyncer.Filename,
			MaxSize:    int(size),
			MaxBackups: cfg.Logger.FileSyncer.MaxBackups,
			MaxAge:     cfg.Logger.FileSyncer.MaxAge,
			Compress:   cfg.Logger.FileSyncer.Compress,
		}
		builder = builder.WithFileConfig(fileConf)
	}

	return builder.Build()
}
