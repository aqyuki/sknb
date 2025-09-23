package logging

import (
	"context"
	"sync"

	"github.com/aqyuki/sknb/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const loggerKey = contextKey("logger")

var (
	defaultLogger     *zap.Logger
	defaultLoggerOnce sync.Once
)

type Level string

const (
	LevelDebug = Level("debug")
	LevelInfo  = Level("info")
	LevelWarn  = Level("warn")
	LevelError = Level("error")
	LevelFatal = Level("fatal")
)

func New() *zap.Logger {
	env := config.Environment(viper.GetString("app.env"))
	return NewForEnv(env)
}

func NewForEnv(env config.Environment) *zap.Logger {
	var conf zap.Config

	if env == config.EnvLocal || env == config.EnvTest {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	level := toZapLevel(Level(viper.GetString("app.log.level")))
	conf.Level = zap.NewAtomicLevelAt(level)

	logger, err := conf.Build()
	if err != nil {
		logger = zap.NewNop()
	}

	return logger
}

func DefaultLogger() *zap.Logger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = New()
	})

	return defaultLogger
}

func toZapLevel(level Level) zapcore.Level {
	switch level {
	case LevelDebug:
		return zap.DebugLevel
	case LevelInfo:
		return zap.InfoLevel
	case LevelWarn:
		return zap.WarnLevel
	case LevelError:
		return zap.ErrorLevel
	case LevelFatal:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}

	return DefaultLogger()
}
