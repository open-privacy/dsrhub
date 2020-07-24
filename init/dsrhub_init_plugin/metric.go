package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func (p *DSRHubInitPlugin) setupMetrics() error {
	if !p.StatsdEnabled || !p.StatsdAPMEnabled {
		return nil
	}

	p.service.Server.WithCustomMiddlewares(
		gintrace.Middleware(p.StatsdAPMServiceName))

	logrus.Info("starting DSRHubInitPlugin tracer metrics...")

	tracer.Start(
		tracer.WithAnalytics(true),
		tracer.WithAgentAddr(fmt.Sprintf("%s:%s", p.StatsdHost, p.StatsdAPMPort)),
	)

	return nil
}
