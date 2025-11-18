package logger

import (
	"go-ecommerce-backend-api/pkg/setting"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {

	logLevel := config.Log_level

	// debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()

	hook := &lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_backups,   // days
		Compress:   config.Compress, // enabled gzip
	}

	writer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(hook),
	)

	core := zapcore.NewCore(encoder, writer, level)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return &LoggerZap{
		Logger: logger,
	}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
