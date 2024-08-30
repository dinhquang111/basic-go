package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.Logger
}

// func newZapLogger() (Logger, error) {
// 	logger, err := zap.NewProduction()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &zapLogger{logger: logger}, nil
// }

func newZapLogger() (Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = ""
	customConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	customCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(customConfig),
		zapcore.AddSync(os.Stdout),
		config.Level,
	)

	// Combine the custom core with the default core
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig), zapcore.AddSync(os.Stdout), config.Level),
		customCore,
	)
	logger := zap.New(core)
	return &zapLogger{logger: logger}, nil
}

func (z *zapLogger) Info(msg string, fields ...Field) {
	z.logger.Info(msg, convertToZapFields(fields)...)
}

func (z *zapLogger) Error(msg string, fields ...Field) {
	z.logger.Error(msg, convertToZapFields(fields)...)
}

func (z *zapLogger) Debug(msg string, fields ...Field) {
	z.logger.Debug(msg, convertToZapFields(fields)...)
}

func convertToZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}
