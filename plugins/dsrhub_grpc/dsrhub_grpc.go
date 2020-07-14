package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dsrhub/dsrhub/idl_dsrhub"
	"github.com/ovh/utask/pkg/plugins/taskplugin"
	"google.golang.org/grpc"
)

var (
	// Plugin is the dsrhub_grpc plugin that we can use as a customized plugin
	// nolint
	Plugin = taskplugin.New(
		"dsrhub_grpc", "1.0", exec,
		taskplugin.WithConfig(validConfig, DsrHubGRPCConfig{}),
	)
)

const (
	defaultTimeoutSeconds = 10
)

// DsrHubGRPCConfig is the configuration needed to perform an gRPC client side call
type DsrHubGRPCConfig struct {
	URL     string                      `json:"url"`
	Request idl_dsrhub.CreateDSRRequest `json:"request"`
	Timeout int                         `json:"timeout,omitempty"` // timeout in seconds
}

func validConfig(config interface{}) error {
	cfg := config.(*DsrHubGRPCConfig)
	if cfg.URL == "" {
		return fmt.Errorf("invalid dsrhub_grpc config url: empty url")
	}
	if cfg.Request.Regulation != "gdpr" && cfg.Request.Regulation != "ccpa" {
		return fmt.Errorf("invalid dsrhub_grpc config request.regulation: %s, want: [gdpr, ccpa]", cfg.Request.Regulation)
	}
	if cfg.Request.SubjectRequestId == "" {
		return fmt.Errorf("invalid dsrhub_grpc config request.subject_request_id: empty request.subject_request_id")
	}
	if cfg.Request.SubjectRequestType == "" {
		return fmt.Errorf("invalid dsrhub_grpc config request.subject_request_type: empty request.subject_request_type")
	}
	if cfg.Request.IdentityValue == "" {
		return fmt.Errorf("invalid dsrhub_grpc config request.identity_value: empty request.identity_value")
	}
	return nil
}

func exec(stepName string, config interface{}, execCtx interface{}) (output interface{}, metadata interface{}, err error) {
	cfg := config.(*DsrHubGRPCConfig)

	// TODO: support secure connection with configuration
	conn, err := grpc.Dial(cfg.URL, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()

	if cfg.Timeout == 0 {
		cfg.Timeout = defaultTimeoutSeconds
	}
	ctx, cancel := context.WithDeadline(
		context.TODO(),
		time.Now().Add(time.Duration(cfg.Timeout)*time.Second),
	)
	defer cancel()

	res, err := idl_dsrhub.NewDSRHubServiceClient(conn).CreateDSR(ctx, &cfg.Request)
	if err != nil {
		return nil, nil, err
	}
	return res, nil, nil
}
