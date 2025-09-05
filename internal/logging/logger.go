// Package logging はログに関係する機能を提供する
package logging

import (
	"context"
	"sync"
	"time"

	"github.com/aqyuki/sknb/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextkey string

const loggerKey = contextkey("logger")

type Level string

const (
	LevelDebug  = Level("debug")
	LevelInfo   = Level("info")
	LevelWarn   = Level("warn")
	LevelError  = Level("error")
	LevelDPanic = Level("dpanic")
	LevelPanic  = Level("panic")
	LevelFatal  = Level("fatal")
)

var (
	defaultLogger     *zap.Logger
	defaultLoggerOnce sync.Once
)

func NewLoggerFromConfig() *zap.Logger {
	level := Level(viper.GetString("app.log"))
	return NewLogger(level)
}

func NewLogger(level Level) *zap.Logger {
	var conf zap.Config

	if config.IsLocal() {
		conf = zap.Config{
			Level:            zap.NewAtomicLevelAt(toZapLevel(level)),
			Development:      true,
			Encoding:         encodingConsole,
			EncoderConfig:    developmentEncoderConfig,
			OutputPaths:      outputStderr,
			ErrorOutputPaths: outputStderr,
		}
	} else {
		conf = zap.Config{
			Level:            zap.NewAtomicLevelAt(toZapLevel(level)),
			Development:      false,
			Encoding:         encodingJSON,
			EncoderConfig:    productionEncoderConfig,
			OutputPaths:      outputStderr,
			ErrorOutputPaths: outputStderr,
		}
	}

	logger, err := conf.Build()
	if err != nil {
		logger = zap.NewNop()
	}

	return logger
}

func DefaultLogger() *zap.Logger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLoggerFromConfig()
	})

	return defaultLogger
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

const (
	timestamp  = "timestamp"
	severity   = "severity"
	logger     = "logger"
	caller     = "caller"
	message    = "message"
	stacktrace = "stacktrace"

	levelDebug     = "DEBUG"
	levelInfo      = "INFO"
	levelWarning   = "WARNING"
	levelError     = "ERROR"
	levelCritical  = "CRITICAL"
	levelAlert     = "ALERT"
	levelEmergency = "EMERGENCY"

	encodingConsole = "console"
	encodingJSON    = "json"
)

var outputStderr = []string{"stderr"}

var productionEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        timestamp,
	LevelKey:       severity,
	NameKey:        logger,
	CallerKey:      caller,
	MessageKey:     message,
	StacktraceKey:  stacktrace,
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    levelEncoder(),
	EncodeTime:     timeEncoder(),
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

var developmentEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "",
	LevelKey:       "L",
	NameKey:        "N",
	CallerKey:      "C",
	FunctionKey:    zapcore.OmitKey,
	MessageKey:     "M",
	StacktraceKey:  "S",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
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
	case LevelDPanic:
		return zap.DPanicLevel
	case LevelPanic:
		return zap.PanicLevel
	case LevelFatal:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func levelEncoder() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString(levelDebug)
		case zapcore.InfoLevel:
			enc.AppendString(levelInfo)
		case zapcore.WarnLevel:
			enc.AppendString(levelWarning)
		case zapcore.ErrorLevel:
			enc.AppendString(levelError)
		case zapcore.DPanicLevel:
			enc.AppendString(levelCritical)
		case zapcore.PanicLevel:
			enc.AppendString(levelAlert)
		case zapcore.FatalLevel:
			enc.AppendString(levelEmergency)
		}
	}
}

func timeEncoder() zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.RFC3339))
	}
}
