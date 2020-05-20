package config

import (
	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

type dbLogger struct {
	logger *zap.Logger
}

func (dbl *dbLogger) Print(args ...interface{}) {
	dbl.logger.Sugar().Info(args...)
}

var Logger *zap.Logger
var DBLogger *dbLogger

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
	DBLogger = &dbLogger{logger: Logger}
}

func Teardown() {
	Logger.Sync()
}
