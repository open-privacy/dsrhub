package main

import (
	"github.com/caarlos0/env"
	"github.com/ovh/utask/pkg/plugins"
)

var Plugin = NewDSRHubInitPlugin() // nolint

type DSRHubInitPlugin struct {
	// Environment variables
	HTTPCallbackBaseURL  string `env:"DSRHUB_HTTP_CALLBACK_BASE_URL" envDefault:"/dsrhub/callback"`
	StatsdEnabled        bool   `env:"STATSD_ENABLED" envDefault:"false"`
	StatsdHost           string `env:"STATSD_HOST" envDefault:"127.0.0.1"`
	StatsdPort           string `env:"STATSD_PORT" envDefault:"8125"`
	StatsdPrefix         string `env:"STATSD_PREFIX" envDefault:"dsrhub."`
	StatsdAPMEnabled     bool   `env:"STATSD_APM_ENABLED" envDefault:"false"`
	StatsdAPMPort        string `env:"STATSD_APM_PORT" envDefault:"8126"`
	StatsdAPMServiceName string `env:"STATSD_APM_SERVICE_NAME" envDefault:"dsrhub"`

	// utask init plugin's Service entrypoint
	service *plugins.Service
}

func NewDSRHubInitPlugin() plugins.InitializerPlugin {
	return &DSRHubInitPlugin{}
}

func (p *DSRHubInitPlugin) Description() string {
	return "DSRHubInitPlugin"
}

func (p *DSRHubInitPlugin) Init(service *plugins.Service) error {
	if err := env.Parse(p); err != nil {
		return err
	}

	p.service = service

	if err := p.setupHTTPCallback(); err != nil {
		return err
	}
	if err := p.setupMetrics(); err != nil {
		return err
	}

	return nil
}
