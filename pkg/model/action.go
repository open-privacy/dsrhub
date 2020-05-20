package model

import (
	"context"
	"errors"
)

const (
	ActionType          = "type"
	ActionTypeSyncHTTP  = "sync_http"
	ActionTypeAsyncHTTP = "async_http"
)

var (
	errInvalidActionType = errors.New("invalid action type")
)

type Actionable interface {
	Act(ctx context.Context, task *Task) error
}

func NewActionable(action string) (Actionable, error) {
	switch action {
	case ActionTypeSyncHTTP:
		return &ActionSyncHTTP{}, nil
	case ActionTypeAsyncHTTP:
		return &ActionAsyncHTTP{}, nil
	default:
		return nil, errInvalidActionType
	}
}

type ActionSyncHTTP struct {
	URL           string
	Method        string
	Headers       map[string]string
	Body          string
	Timeout       int // timeout in seconds
	AllowRedirect bool
}

func (a *ActionSyncHTTP) Act(ctx context.Context, task *Task) error {
	return nil
}

type ActionAsyncHTTP struct {
	URL           string
	Method        string
	Headers       map[string]string
	Body          string
	Timeout       int // timeout in seconds
	AllowRedirect bool
}

func (a *ActionAsyncHTTP) Act(ctx context.Context, task *Task) error {
	return nil
}
