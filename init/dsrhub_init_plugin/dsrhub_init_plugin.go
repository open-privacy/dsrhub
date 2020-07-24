package main

import (
	"github.com/caarlos0/env"
	"github.com/ovh/utask/pkg/plugins"
)

// Plugin is exposed go plugin variable that we can use in uTask
var Plugin = NewDSRHubInitPlugin() // nolint

// DSRHubInitPlugin is a plugin config struct.
type DSRHubInitPlugin struct {
	// Environment variables
	HTTPCallbackBaseURL  string `env:"DSRHUB_HTTP_CALLBACK_BASE_URL" envDefault:"/dsrhub/callback"`
	StatsdEnabled        bool   `env:"STATSD_ENABLED" envDefault:"false"`
	StatsdHost           string `env:"STATSD_HOST" envDefault:"127.0.0.1"`
	StatsdPort           string `env:"STATSD_PORT" envDefault:"8125"`
	StatsdPrefix         string `env:"STATSD_PREFIX" envDefault:"dsrhub."`
	StatsdAPMEnabled     bool   `env:"STATSD_APM_ENABLED" envDefault:"false"`
	StatsdAPMPort        string `env:"STATSD_APM_PORT" envDefault:"8126"`
	StatsdAPMServiceName string `env:"DD_SERVICE" envDefault:"dsrhub"`

	// utask init plugin's Service entrypoint
	service *plugins.Service
}

// NewDSRHubInitPlugin creates a new DSRHubInitPlugin
func NewDSRHubInitPlugin() plugins.InitializerPlugin {
	return &DSRHubInitPlugin{}
}

// Description is the description
func (p *DSRHubInitPlugin) Description() string {
	return "DSRHubInitPlugin"
}

// Init setup the plugin by extending the official plugins.Service
func (p *DSRHubInitPlugin) Init(service *plugins.Service) error {
	if err := env.Parse(p); err != nil {
		return err
	}

	p.service = service

	if err := p.setupMetrics(); err != nil {
		return err
	}

	if err := p.setupHTTPCallback(); err != nil {
		return err
	}
	return nil
}
