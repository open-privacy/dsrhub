package main

import (
	"context"
	"fmt"

	"github.com/wI2L/fizz"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func (p *DSRHubInitPlugin) setupMetrics() error {
	if !p.StatsdEnabled || !p.StatsdAPMEnabled {
		return nil
	}

	router, ok := p.service.Server.Handler(context.Background()).(*fizz.Fizz)
	if !ok {
		return fmt.Errorf("failed to load router in plugin: %s", p.Description())
	}

	tracer.Start(
		tracer.WithAgentAddr(fmt.Sprintf("%s:%s", p.StatsdHost, p.StatsdAPMPort)),
		tracer.WithService(p.StatsdAPMServiceName),
	)

	ginEngine := router.Engine()
	ginEngine.Use(gintrace.Middleware(
		fmt.Sprintf("%s-http-server", p.StatsdAPMServiceName),
	))

	return nil
}
