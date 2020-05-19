package model

import (
	"context"
	"fmt"
	"time"

	"github.com/dsrhub/dsrhub/pkg/config"
	"github.com/qor/transition"
	"go.uber.org/zap"
)

type Task struct {
	BaseModel
	transition.Transition

	WorkflowID string `gorm:"type:char(36)"`
	Workflow   *Workflow
	Requires   []*Task `gorm:"many2many:task_requires;association_jointable_foreignkey:require_id"`

	Name        string
	Description string
	Input       EncryptionKV `gorm:"type:varchar(65532)"`
	Output      EncryptionKV `gorm:"type:varchar(1000000)"`
	Action      KV           `gorm:"type:varchar(65532)"`
	Publish     KV           `gorm:"type:varchar(65532)"`
	Retry       Retry        `gorm:"type:varchar(65532)"`
	ScopeID     string
	Timeout     int // timeout in seconds

	// gorm ignored
	FSM         *transition.StateMachine `gorm:"-"`
	RequiresMap map[string]*Task         `gorm:"-"`

	// private
	actionable Actionable `gorm:"-"`
}

const (
	TASK_STATE_CREATED   = "created"
	TASK_STATE_COMPLETED = "completed"
	TASK_STATE_CANCELLED = "cancelled"
)

func (t *Task) Prepare() error {
	// prepare fsm
	fsm := transition.New(&Task{})
	fsm.State(TASK_STATE_CREATED)
	fsm.State(TASK_STATE_COMPLETED)
	fsm.State(TASK_STATE_CANCELLED)
	fsm.Initial(TASK_STATE_CREATED)
	t.FSM = fsm

	// prepare actionable
	actionable, err := NewActionable(t.Action)
	if err != nil {
		return err
	}
	t.actionable = actionable
	return nil
}

func (t *Task) Run() error {
	if t.actionable == nil {
		return fmt.Errorf("empty actionable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t.Timeout)*time.Second)
	defer cancel()
	if err := t.actionable.Act(ctx, t); err != nil {
		return err
	}
	return nil
}

func (t *Task) Logger() *zap.Logger {
	return config.Logger.With(
		zap.String("workflow_id", t.WorkflowID),
		zap.String("activity_id", t.ID),
	)
}
