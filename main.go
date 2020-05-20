package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dsrhub/dsrhub/pkg/adminserver"
	"github.com/dsrhub/dsrhub/pkg/config"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		config.Logger.Fatal("failed to run dsrhub", zap.Error(err))
	}
}

func waitForSignal(teardown func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT, syscall.SIGTERM,
	)
	<-signalChan
	teardown()
}

func run() error {
	adminserver := &adminserver.AdminServer{
		HostPort: config.ENV.AdminServerHostPort,
	}
	go adminserver.Start()

	waitForSignal(func() {
		adminserver.Teardown()
		config.Teardown()
	})
	return nil
}
