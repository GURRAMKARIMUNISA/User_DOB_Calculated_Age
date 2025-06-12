package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(env string) {
	var err error
	if env == "production" {
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		Log, err = config.Build()
	} else {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		Log, err = config.Build()
	}

	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(Log) // Set global logger
}

func SyncLogger() {
	if Log != nil {
		Log.Sync()
	}
}