package main

import (
	"context"
	"fmt"

	"github.com/ovh/utask/pkg/plugins"
)

var Plugin = &WebhookInitPlugin{}

var _ plugins.InitializerPlugin = (*WebhookInitPlugin)(nil)

type WebhookInitPlugin struct{}

func (plugin *WebhookInitPlugin) Init(service *plugins.Service) error {
	server := service.Server
	handler := server.Handler(context.TODO())
	fmt.Println("testing from WebhookInitPlugin")
	fmt.Printf("%T", handler)
	return nil
}

func (plugin *WebhookInitPlugin) Description() string {
	return "webhook init plugin"
}
