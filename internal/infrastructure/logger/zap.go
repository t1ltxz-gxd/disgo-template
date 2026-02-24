package logger

import (
	"errors"
	"strings"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	envDev         = "dev"
	envDevelop     = "develop"
	envDevelopment = "development"
	envProd        = "prd"
	envProduction  = "production"
)

// ZapLoggerBuilder is used to create and configure the zap.Logger.
type ZapLoggerBuilder struct {
	env      string
	name     string
	fileConf *lumberjack.Logger
}

// NewZapLoggerBuilder creates a new instance of ZapLoggerBuilder to configure the logger.
func NewZapLoggerBuilder() *ZapLoggerBuilder {
	return &ZapLoggerBuilder{}
}

// WithEnv sets the environment in which the logger will work.
func (b *ZapLoggerBuilder) WithEnv(env string) *ZapLoggerBuilder {
	b.env = strings.ToLower(env)
	return b
}

// WithName sets the name for the logger to be used as a prefix in the logs.
func (b *ZapLoggerBuilder) WithName(name string) *ZapLoggerBuilder {
	b.name = name
	return b
}

// WithFileConfig sets the configuration to write logs to the file.
func (b *ZapLoggerBuilder) WithFileConfig(fileConf *lumberjack.Logger) *ZapLoggerBuilder {
	b.fileConf = fileConf
	return b
}

// Build creates and returns a new instance zap.Logger depending on the given environment.
func (b *ZapLoggerBuilder) Build() (*zap.Logger, error) {
	var zapLog *zap.Logger
	var ws zapcore.WriteSyncer
	var encoder zapcore.Encoder

	switch b.env {
	case envDev, envDevelopment, envDevelop:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		// Adding a call skip to work correctly with alias.go
		zapLog, _ = config.Build(zap.AddCallerSkip(1))
	case envProd, envProduction:
		if b.fileConf == nil {
			return nil, errors.New("file config required for production logger")
		}
		fileSyncer := zapcore.AddSync(b.fileConf)
		consoleSyncer := zapcore.AddSync(colorable.NewColorableStdout())
		ws = zapcore.NewMultiWriteSyncer(fileSyncer, consoleSyncer)
		config := zap.NewProductionConfig()
		encoder = zapcore.NewJSONEncoder(config.EncoderConfig)
		core := zapcore.NewCore(encoder, ws, config.Level)
		// Adding a call skip to work correctly with alias.go
		zapLog = zap.New(core, zap.AddCallerSkip(1))
	default:
		err := errors.New("unknown environment")
		config := zap.NewDevelopmentConfig()
		// Adding a call skip to work correctly with alias.go
		zapLog, _ = config.Build(zap.AddCallerSkip(1))
		zapLog.Warn("Failed to initialize logger with proper environment", zap.Error(err), zap.String("env", b.env))
	}
	if b.name != "" {
		zapLog = zapLog.Named(b.name)
	}
	zap.ReplaceGlobals(zapLog)
	return zapLog, nil
}

// GetLogLevelByMode returns the level of logging depending on the given environment.
func GetLogLevelByMode(env string) zapcore.Level {
	switch strings.ToLower(env) {
	case "production", "prd", "prod":
		return zap.InfoLevel
	case "test":
		return zap.WarnLevel
	case "dev", "development":
		return zap.DebugLevel
	default:
		return zap.DebugLevel
	}
}
