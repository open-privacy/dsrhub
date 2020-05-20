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

	Name        string
	Description string
	Timeout     int    // timeout in seconds
	Action      string `gorm:"type:varchar(65532)"`
	Input       KV     `gorm:"type:varchar(65532)"`
	Publish     KV     `gorm:"type:varchar(65532)"`
	Retry       Retry  `gorm:"type:varchar(65532)"`

	// -------- yaml ignored --------------------------------------------------
	WorkflowID    string       `gorm:"type:char(36)" yaml:"-"`
	Workflow      *Workflow    `yaml:"-"`
	Output        EncryptionKV `gorm:"type:varchar(1000000)" yaml:"-"`
	RenderedInput EncryptionKV `gorm:"type:varchar(65532)" yaml:"-"`
	RequiresList  []*Task      `gorm:"many2many:task_requires;association_jointable_foreignkey:require_task_id" yaml:"-"`
	ScopeID       string       `gorm:"index" yaml:"-"`
	// ------------------------------------------------------------------------

	// -------- gorm ignored --------------------------------------------------
	Requires []string `gorm:"-"` // list of the task names that the task requires
	// ------------------------------------------------------------------------

	// ---------gorm and yaml ignored -----------------------------------------
	FSM *transition.StateMachine `gorm:"-" yaml:"-"`
	// ------------------------------------------------------------------------

	// private
	actionable Actionable `gorm:"-"`
}

const (
	TASK_STATE_CREATED   = "@created"
	TASK_STATE_COMPLETED = "@completed"
	TASK_STATE_FAILED    = "@failed"
	TASK_STATE_CANCELLED = "@cancelled"
)

func (t *Task) Prepare() error {
	// prepare fsm
	fsm := transition.New(&Task{})
	fsm.State(TASK_STATE_CREATED)
	fsm.State(TASK_STATE_COMPLETED)
	fsm.State(TASK_STATE_FAILED)
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
