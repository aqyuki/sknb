package logging

import (
	"testing"

	"github.com/aqyuki/sknb/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func Test_toZapLevel(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		level    string
		expected zapcore.Level
	}{
		"debug":   {level: "debug", expected: zapcore.DebugLevel},
		"info":    {level: "info", expected: zapcore.InfoLevel},
		"warn":    {level: "warn", expected: zapcore.WarnLevel},
		"error":   {level: "error", expected: zapcore.ErrorLevel},
		"fatal":   {level: "fatal", expected: zapcore.FatalLevel},
		"unknown": {level: "unknown", expected: zapcore.InfoLevel},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := toZapLevel(Level(test.level))
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestWithLogger(t *testing.T) {
	t.Parallel()

	logger := NewForEnv(config.EnvTest)
	ctx := WithLogger(t.Context(), logger)

	require.NotNil(t, ctx)
}

func TestFromContext(t *testing.T) {
	t.Parallel()

	t.Run("return given logger", func(t *testing.T) {
		t.Parallel()

		logger := NewForEnv(config.EnvTest)
		ctx := WithLogger(t.Context(), logger)

		actual := FromContext(ctx)
		require.NotNil(t, actual)
		require.Equal(t, logger, actual)
	})

	t.Run("return default logger", func(t *testing.T) {
		t.Parallel()

		actual := FromContext(t.Context())
		require.NotNil(t, actual)
		require.Equal(t, defaultLogger, actual)
	})
}
