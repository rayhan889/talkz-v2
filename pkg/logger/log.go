package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger(env string) {
	var c zap.Config

	if env == "production" {
		c = zap.NewProductionConfig()
		c.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	} else if env == "development" {
		c = zap.NewDevelopmentConfig()
		c.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	l, err := c.Build()
	if err != nil {
		panic(err)
	}

	Log = l.Sugar()
}
