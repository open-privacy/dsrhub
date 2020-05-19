package config

import (
	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	setupEnv()
	setupLogger()
}

func setupEnv() {
	env.Parse(&ENV)
}

func setupLogger() {
	var l *zap.Logger
	if ENV.LoggerEnv == "production" {
		l, _ = zap.NewProduction()
	} else {
		l, _ = zap.NewDevelopment()
	}
	Logger = l
}

func Teardown() {
	Logger.Sync()
}
