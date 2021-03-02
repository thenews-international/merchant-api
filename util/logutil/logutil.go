package logutil

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	consoleEncoding = "console"
	jsonEncoding    = "json"
)

func NewLogger(filters string, format string, logFile string) (*zap.Logger, func(), error) {
	if filters == "" {
		cleanup := func() {}
		return zap.NewNop(), cleanup, nil
	}

	stableWidthNameEncoder := func(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("%-18s", loggerName))
	}
	stableWidthCapitalLevelEncoder := func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("%-5s", l.CapitalString()))
	}
	const (
		Black uint8 = iota + 30
		Red
		Green
		Yellow
		Blue
		Magenta
		Cyan
		White
	)
	stableWidthCapitalColorLevelEncoder := func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Magenta, "DEBUG"))
		case zapcore.InfoLevel:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Blue, "INFO "))
		case zapcore.WarnLevel:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Yellow, "WARN "))
		case zapcore.ErrorLevel:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Red, "ERROR"))
		case zapcore.DPanicLevel:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Red, "DPANIC"))
		case zapcore.PanicLevel:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Red, "PANIC"))
		case zapcore.FatalLevel:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Red, "FATAL"))
		default:
			enc.AppendString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", Red, l.CapitalString()))
		}
	}

	// configure zap
	var config zap.Config
	switch strings.ToLower(format) {
	case "":
		config = zap.NewDevelopmentConfig()
	case "json":
		config = zap.NewProductionConfig()
		config.Development = true
		config.Encoding = jsonEncoding
	case "light-json":
		config = zap.NewProductionConfig()
		config.Encoding = jsonEncoding
		config.EncoderConfig.TimeKey = ""
		config.EncoderConfig.EncodeLevel = stableWidthCapitalLevelEncoder
		config.Development = true
		config.DisableStacktrace = true
	case "light-console":
		config = zap.NewDevelopmentConfig()
		config.Encoding = consoleEncoding
		config.EncoderConfig.TimeKey = ""
		config.EncoderConfig.EncodeLevel = stableWidthCapitalLevelEncoder
		config.DisableStacktrace = true
		config.EncoderConfig.EncodeName = stableWidthNameEncoder
		config.Development = true
	case "light-color":
		config = zap.NewDevelopmentConfig()
		config.Encoding = consoleEncoding
		config.EncoderConfig.TimeKey = ""
		config.EncoderConfig.EncodeLevel = stableWidthCapitalColorLevelEncoder
		config.DisableStacktrace = true
		config.EncoderConfig.EncodeName = stableWidthNameEncoder
		config.Development = true
	case "console":
		config = zap.NewDevelopmentConfig()
		config.Encoding = consoleEncoding
		config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		config.EncoderConfig.EncodeLevel = stableWidthCapitalLevelEncoder
		config.DisableStacktrace = true
		config.EncoderConfig.EncodeName = stableWidthNameEncoder
		config.Development = true
	case "color":
		config = zap.NewDevelopmentConfig()
		config.Encoding = consoleEncoding
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
		config.EncoderConfig.EncodeLevel = stableWidthCapitalColorLevelEncoder
		config.DisableStacktrace = true
		config.EncoderConfig.EncodeName = stableWidthNameEncoder
		config.Development = true
	default:
		return nil, nil, fmt.Errorf("unknown log format: %q", format)
	}

	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	switch logFile {
	case "":
	case "stdout", "stderr":
		config.OutputPaths = []string{logFile}
	default:
		config.OutputPaths = []string{logFile}
	}

	base, err := config.Build()
	if err != nil {
		return nil, nil, err
	}

	return DecorateLogger(base)
}

func DecorateLogger(base *zap.Logger) (*zap.Logger, func(), error) {
	logger := zap.New(base.Core(), zap.AddCaller())
	zap.ReplaceGlobals(logger.Named("other"))

	cleanup := func() {
		_ = logger.Sync()
	}

	return logger.Named("bty"), cleanup, nil
}
