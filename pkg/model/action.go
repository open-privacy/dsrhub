package model

import (
	"context"
	"errors"

	"github.com/spf13/cast"
)

const (
	ACTION_TYPE            = "type"
	ACTION_TYPE_SYNC_HTTP  = "sync_http"
	ACTION_TYPE_ASYNC_HTTP = "async_http"
)

var (
	errEmptyAction       = errors.New("empty action")
	errEmptyActionType   = errors.New("empty action type")
	errInvalidActionType = errors.New("invalid action type")
)

type Actionable interface {
	Act(ctx context.Context, task *Task) error
}

func NewActionable(action KV) (Actionable, error) {
	if action == nil {
		return nil, errEmptyAction
	}
	v, ok := action[ACTION_TYPE]
	if !ok {
		return nil, errEmptyActionType
	}
	aType := cast.ToString(v)
	switch aType {
	case ACTION_TYPE_SYNC_HTTP:
		return &ActionSyncHTTP{}, nil
	case ACTION_TYPE_ASYNC_HTTP:
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
